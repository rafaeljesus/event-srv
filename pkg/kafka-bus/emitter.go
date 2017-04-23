package kafkabus

import (
	"encoding/json"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
)

type Emitter interface {
	Emit() chan *Message
	Close()
}

type Message struct {
	Topic     string
	Payload   interface{}
	Partition int32
}

type EmitterConfig struct {
	Url      string
	Attempts int
	Timeout  time.Duration
}

type kafkaEmitter struct {
	emitter     sarama.AsyncProducer
	emitterChan chan *Message
}

func NewEmitter(c EmitterConfig) (emitter Emitter, err error) {
	brokers := []string{c.Url}
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Producer.Retry.Max = c.Attempts
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewAsyncProducer(brokers, nil)
	if err != nil {
		return
	}

	ke := &kafkaEmitter{
		emitter:     producer,
		emitterChan: make(chan *Message),
	}

	go ke.register()

	emitter = ke

	return
}

func (ke *kafkaEmitter) Emit() chan *Message {
	return ke.emitterChan
}

func (ke *kafkaEmitter) register() {
	for {
		select {
		case m := <-ke.emitterChan:
			ke.emit(m)
		}
	}
}

func (ke *kafkaEmitter) emit(m *Message) (err error) {
	l := log.WithFields(log.Fields{
		"topic":     m.Topic,
		"partition": m.Partition,
	})

	l.Debug("Sending event")

	p, err := json.Marshal(m.Payload)
	if err != nil {
		return
	}

	message := &sarama.ProducerMessage{
		Topic:     m.Topic,
		Partition: m.Partition,
		Value:     sarama.StringEncoder(p),
	}

	for {
		select {
		case ke.emitter.Input() <- message:
		case <-ke.emitter.Successes():
			l.Info("Event successfully sent")
			return
		case errors := <-ke.emitter.Errors():
			l.WithError(err).Error("Failed to send event")
			err = errors
			return
		}
	}
}

func (ke *kafkaEmitter) Close() {
	ke.emitter.Close()
}
