package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pressly/chi"
	"github.com/rafaeljesus/event-srv/pkg/mocks"
	"github.com/rafaeljesus/event-srv/pkg/models"
)

func TestEventsIndex(t *testing.T) {
	repoMock := mocks.NewEventRepo()
	emitterMock := mocks.NewEmitter()
	h := NewEventsHandler(repoMock, emitterMock)

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/events", nil)
	if err != nil {
		t.Fail()
	}

	r := chi.NewRouter()
	r.Get("/events", h.EventsIndex)
	r.ServeHTTP(res, req)

	events := []models.Event{}
	if err := json.NewDecoder(res.Body).Decode(&events); err != nil {
		t.Fail()
	}

	if len(events) == 0 {
		t.Fail()
	}

	if res.Code != http.StatusOK {
		t.Errorf("Expected status %d to be equal %d", res.Code, http.StatusOK)
	}
}

func TestEventsIndexByUUID(t *testing.T) {
	repoMock := mocks.NewEventRepo()
	h := NewEventsHandler(repoMock, nil)

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/events?uuid=12kh312uynb2u", nil)
	if err != nil {
		t.Fail()
	}

	r := chi.NewRouter()
	r.Get("/events", h.EventsIndex)
	r.ServeHTTP(res, req)

	events := []models.Event{}
	if err := json.NewDecoder(res.Body).Decode(&events); err != nil {
		t.Fail()
	}

	if len(events) == 0 {
		t.Fail()
	}

	if !repoMock.ByUUID {
		t.Fail()
	}

	if res.Code != http.StatusOK {
		t.Errorf("Expected status %d to be equal %d", res.Code, http.StatusOK)
	}
}

func TestEventsIndexByName(t *testing.T) {
	repoMock := mocks.NewEventRepo()
	h := NewEventsHandler(repoMock, nil)

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/events?name=something_happened", nil)
	if err != nil {
		t.Fail()
	}

	r := chi.NewRouter()
	r.Get("/events", h.EventsIndex)
	r.ServeHTTP(res, req)

	events := []models.Event{}
	if err := json.NewDecoder(res.Body).Decode(&events); err != nil {
		t.Fail()
	}

	if len(events) == 0 {
		t.Fail()
	}

	if !repoMock.ByName {
		t.Fail()
	}

	if res.Code != http.StatusOK {
		t.Errorf("Expected status %d to be equal %d", res.Code, http.StatusOK)
	}
}

func TestEventsIndexByStatus(t *testing.T) {
	repoMock := mocks.NewEventRepo()
	h := NewEventsHandler(repoMock, nil)

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/events?status=something_processed", nil)
	if err != nil {
		t.Fail()
	}

	r := chi.NewRouter()
	r.Get("/events", h.EventsIndex)
	r.ServeHTTP(res, req)

	events := []models.Event{}
	if err := json.NewDecoder(res.Body).Decode(&events); err != nil {
		t.Fail()
	}

	if len(events) == 0 {
		t.Fail()
	}

	if !repoMock.ByStatus {
		t.Fail()
	}

	if res.Code != http.StatusOK {
		t.Errorf("Expected status %d to be equal %d", res.Code, http.StatusOK)
	}
}

func TestEventsCreate(t *testing.T) {
	repoMock := mocks.NewEventRepo()
	emitterMock := mocks.NewEmitter()
	h := NewEventsHandler(repoMock, emitterMock)

	res := httptest.NewRecorder()
	body := strings.NewReader(`{"name":"foo","status":"bar","payload":"some_data"}`)
	req, err := http.NewRequest("POST", "/events", body)
	if err != nil {
		t.Fail()
	}

	r := chi.NewRouter()
	r.Post("/events", h.EventsCreate)
	r.ServeHTTP(res, req)

	if !repoMock.Created {
		t.Fail()
	}

	if !emitterMock.Emitted {
		t.Fail()
	}

	if res.Code != http.StatusAccepted {
		t.Errorf("Expected status %d to be equal %d", res.Code, http.StatusAccepted)
	}
}
