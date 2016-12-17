package handlers

import (
	"github.com/rafaeljesus/event-srv/event_bus"
	"github.com/rafaeljesus/event-srv/models"
)

type Env struct {
	Repo     models.Repo
	EventBus *event_bus.EventBus
}
