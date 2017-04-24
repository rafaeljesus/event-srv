package kafkabus

type Message struct {
	Topic     string
	Payload   interface{}
	Partition int32
}
