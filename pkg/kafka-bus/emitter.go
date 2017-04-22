package kafkabus

import (
	"encoding/json"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
)

type Emitter interface {
	Emit(topic string, payload interface{}) error
	Close()
}

type EmitterConfig struct {
	Url      string
	Attempts int
	Timeout  time.Duration
}

type kafkaEmitter struct {
	Emitter sarama.AsyncProducer
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

	emitter = &kafkaEmitter{producer}

	return
}

func (e *kafkaEmitter) Emit(topic string, payload interface{}) (err error) {
	partition := 1<<32 - 1

	l := log.WithFields(log.Fields{
		"topic":     topic,
		"partition": partition,
	})

	l.Debug("Sending event")

	p, err := json.Marshal(payload)
	if err != nil {
		return
	}

	message := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: partition,
		Value:     sarama.StringEncoder(p),
	}

	for {
		select {
		case e.Emitter.Input() <- message:
		case <-e.Emitter.Successes():
			l.Info("Event successfully sent")
			return
		case errors := <-e.Emitter.Errors():
			l.WithError(err).Error("Failed to send event")
			err = errors
			return
		}
	}
}

func (e *kafkaEmitter) Close() {
	e.Emitter.Close()
}
