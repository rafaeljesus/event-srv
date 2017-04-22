package handlers

import (
	"net/http"

	"github.com/rafaeljesus/event-srv/pkg/models"
)

func (env *Env) EventsIndex(c echo.Context) error {
	cid := c.QueryParam("cid")
	name := c.QueryParam("name")
	status := c.QueryParam("status")

	query := models.NewQuery(cid, name, status)
	events := []models.Event{}
	if err := env.Repo.SearchEvents(query, &events); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, events)
}

func (env *Env) EventsCreate(c echo.Context) error {
	event := &models.Event{}
	if err := c.Bind(event); err != nil {
		return err
	}

	if err := env.EventBus.Emit("events", event); err != nil {
		return err
	}

	response := map[string]string{"ok": "true"}

	return c.JSON(http.StatusAccepted, response)
}
