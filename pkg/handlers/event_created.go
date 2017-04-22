package handlers

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/rafaeljesus/event-srv/pkg/models"
)

func (env *Env) EventCreated(message []byte) {
	event := &models.Event{}
	if err := json.Unmarshal(message, &event); err != nil {
		log.WithError(err).Error("Failed to decode message!")
	}

	if err := env.Repo.CreateEvent(event); err != nil {
		log.WithError(err).Error("Failed to insert event!")
	}
}
