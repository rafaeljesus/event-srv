package handlers

import (
	"github.com/rafaeljesus/trace-srv/event_bus"
	"github.com/rafaeljesus/trace-srv/models"
)

type Env struct {
	Repo     models.Repo
	EventBus *event_bus.EventBus
}
