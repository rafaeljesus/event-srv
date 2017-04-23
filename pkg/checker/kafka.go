package checker

import (
	"github.com/Shopify/sarama"
)

type kafka struct {
	url string
}

func NewKafka(url string) *kafka {
	return &kafka{url}
}

func (k *kafka) IsAlive() bool {
	brokers := []string{k.url}
	config := sarama.NewConfig()

	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		return false
	}

	defer consumer.Close()

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return false
	}

	defer producer.Close()

	return true
}
