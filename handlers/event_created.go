package handlers

import (
	"encoding/json"
	"github.com/rafaeljesus/trace-srv/models"
	"log"
)

func (env *Env) EventCreated(message []byte) {
	event := &models.Event{}
	if err := json.Unmarshal(message, &event); err != nil {
		log.Fatalln("Failed to parse message", err)
	}

	if err := env.Repo.CreateEvent(event); err != nil {
		log.Fatalln("Failed to insert event", err)
	}
}
