package mocks

import (
	"github.com/rafaeljesus/event-srv/pkg/models"
)

type EventRepoMock struct {
	Created      bool
	Searched     bool
	ByStatus     bool
	ByExpression bool
}

func NewEventRepo() *EventRepoMock {
	return &EventRepoMock{
		Created:      false,
		Searched:     false,
		ByStatus:     false,
		ByExpression: false,
	}
}

func (repo *EventRepoMock) Create(event *models.Event) (err error) {
	repo.Created = true
	return
}

func (repo *EventRepoMock) Search(sc *models.Query) (events []models.Event, err error) {
	events = append(events, models.Event{Expression: "* * * * * *"})

	switch true {
	case sc.Status != "":
		repo.ByStatus = true
	case sc.Expression != "":
		repo.ByExpression = true
	default:
		repo.Searched = true
	}

	return
}
