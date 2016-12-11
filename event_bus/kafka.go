package event_bus

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
)

type EventBus struct {
	Emitter  sarama.SyncProducer
	Listener sarama.Consumer
}

func NewEventBus(url string) (*EventBus, error) {
	brokers := []string{url}
	consumer, _ := sarama.NewConsumer(brokers, nil)

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll

	producer, _ := sarama.NewSyncProducer(brokers, config)

	return &EventBus{producer, consumer}, nil
}

func (bus *EventBus) Emit(topic string, payload interface{}) error {
	msg, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	message := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(msg),
	}

	if _, _, err := bus.Emitter.SendMessage(message); err != nil {
		log.WithError(err).Error("Failed to emit message!")
		return err
	}

	return nil
}

func (bus *EventBus) On(topic string, fn func([]byte)) error {
	pc, err := bus.Listener.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		return err
	}

	go func() {
		for message := range pc.Messages() {
			fn(message.Value)
		}
	}()

	return nil
}
