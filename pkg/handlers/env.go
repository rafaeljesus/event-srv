package handlers

import (
	"github.com/rafaeljesus/event-srv/pkg/kafka-bus"
	"github.com/rafaeljesus/event-srv/pkg/models"
)

type Env struct {
	Repo     models.Repo
	EventBus eventBus.Bus
}
