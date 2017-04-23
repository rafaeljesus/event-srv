package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rafaeljesus/event-srv/pkg/kafka-bus"
	"github.com/rafaeljesus/event-srv/pkg/models"
	"github.com/rafaeljesus/event-srv/pkg/render"
	"github.com/rafaeljesus/event-srv/pkg/repos"
)

type EventsHandler struct {
	EventRepo repos.EventRepo
	Emitter   kafkabus.Emitter
}

func NewEventsHandler(r repos.EventRepo, e kafkabus.Emitter) *EventsHandler {
	return &EventsHandler{r, e}
}

func (h *EventsHandler) EventsIndex(w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("uuid")
	name := r.URL.Query().Get("name")
	status := r.URL.Query().Get("status")

	query := models.NewQuery(uuid, name, status)
	events, err := h.EventRepo.Find(query)
	if err != nil {
		render.JSON(w, http.StatusPreconditionFailed, err)
	}

	render.JSON(w, http.StatusOK, events)
}

func (h *EventsHandler) EventsCreate(w http.ResponseWriter, r *http.Request) {
	event := new(models.Event)
	if err := json.NewDecoder(r.Body).Decode(event); err != nil {
		render.JSON(w, http.StatusBadRequest, "Failed to decode request body")
		return
	}

	h.Emitter.Emit() <- &kafkabus.Message{
		Topic:     "events",
		Payload:   event,
		Partition: -1,
	}

	render.JSON(w, http.StatusCreated, "OK")
}
