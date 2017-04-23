package mocks

import (
	"github.com/rafaeljesus/event-srv/pkg/models"
)

type EventRepoMock struct {
	Created  bool
	Found    bool
	ByUUID   bool
	ByName   bool
	ByStatus bool
}

func NewEventRepo() *EventRepoMock {
	return &EventRepoMock{
		Created:  false,
		Found:    false,
		ByUUID:   false,
		ByName:   false,
		ByStatus: false,
	}
}

func (repo *EventRepoMock) Create(event *models.Event) (err error) {
	repo.Created = true
	return
}

func (repo *EventRepoMock) Find(sc *models.Query) (events []models.Event, err error) {
	events = append(events, models.Event{
		UUID:   "12kh312uynb2u",
		Name:   "something_happened",
		Status: "something_processed",
	})

	switch true {
	case sc.UUID != "":
		repo.ByUUID = true
		return
	case sc.Name != "":
		repo.ByName = true
		return
	case sc.Status != "":
		repo.ByStatus = true
		return
	default:
		repo.Found = true
		return
	}
}
