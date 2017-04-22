package handlers

import (
	"net/http"

	"github.com/rafaeljesus/event-srv/pkg/checker"
	"github.com/rafaeljesus/event-srv/pkg/render"
)

type HealthzHandler struct {
	checkers map[string]checker.Checker
}

func NewHealthzHandler(checkers map[string]checker.Checker) *HealthzHandler {
	return &HealthzHandler{checkers}
}

func (h *HealthzHandler) HealthzIndex(w http.ResponseWriter, r *http.Request) {
	payload := make(map[string]bool)

	for k, v := range h.checkers {
		payload[k] = v.IsAlive()
	}

	render.JSON(w, http.StatusOK, payload)
}
