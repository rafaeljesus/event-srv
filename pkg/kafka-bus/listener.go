package kafkabus

import (
	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
)

type Listener interface {
	On(topic string, partition int32, fn fnHandler) error
	Close()
}

type kafkaListener struct {
	listener sarama.Consumer
}

type fnHandler func(payload []byte) error

func NewListener(c Config) (listener Listener, err error) {
	brokers := []string{c.Url}
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return
	}

	listener = &kafkaListener{consumer}

	return
}

func (kl *kafkaListener) On(topic string, partition int32, fn fnHandler) (err error) {
	pc, err := kl.listener.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		return
	}

	go func() {
		for {
			select {
			case message := <-pc.Messages():
				log.WithFields(log.Fields{
					"topic":     topic,
					"partition": partition,
				}).Debug("receiving a message from broker")

				fn(message.Value)
			}
		}
	}()

	return
}

func (kl *kafkaListener) Close() {
	kl.listener.Close()
}
