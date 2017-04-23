package mocks

import (
	"github.com/rafaeljesus/event-srv/pkg/kafka-bus"
)

type EmitterMock struct {
	Emitted bool
	c       chan *kafkabus.Message
}

func NewEmitter() *EmitterMock {
	return &EmitterMock{
		c: make(chan *kafkabus.Message),
	}
}

func (e *EmitterMock) Emit() chan *kafkabus.Message {
	e.Emitted = true
	go func() {
		<-e.c
	}()
	return e.c
}

func (e *EmitterMock) Close() {}
