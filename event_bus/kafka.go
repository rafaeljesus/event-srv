package event_bus

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
)

type EventBus struct {
	Emitter  sarama.AsyncProducer
	Listener sarama.Consumer
}

func NewEventBus(url string) (*EventBus, error) {
	brokers := []string{url}
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Producer.Retry.Max = 5
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		return nil, err
	}

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &EventBus{producer, consumer}, nil
}

func (bus *EventBus) Emit(topic string, payload interface{}) error {
	msg, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}

	for {
		select {
		case bus.Emitter.Input() <- message:
			return nil
		case <-bus.Emitter.Successes():
			log.WithField("topic", topic).Info("Message emitted")
		case err := <-bus.Emitter.Errors():
			log.WithError(err).Error("Failed to emit message!")
			return err
		}
	}
}

func (bus *EventBus) On(topic string, fn func([]byte)) error {
	partitions, err := bus.Listener.Partitions(topic)
	if err != nil {
		return err
	}

	for _, partition := range partitions {
		pc, err := bus.Listener.ConsumePartition(topic, partition, sarama.OffsetOldest)
		if err != nil {
			return err
		}

		go func(pc sarama.PartitionConsumer) {
			for {
				select {
				case message := <-pc.Messages():
					log.WithField("topic", topic).Info("Message consumed")
					fn(message.Value)
				case err := <-pc.Errors():
					log.WithError(err).Error("Failed to consume message!")
				}
			}
		}(pc)
	}

	return nil
}
