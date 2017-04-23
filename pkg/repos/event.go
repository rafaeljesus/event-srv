package repos

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/rafaeljesus/event-srv/pkg/models"
	"gopkg.in/olivere/elastic.v5"
)

type EventRepo interface {
	Create(e *models.Event) error
	Find(q *models.Query) ([]models.Event, error)
}

type Event struct {
	db *elastic.Client
}

func NewEvent(db *elastic.Client) *Event {
	return &Event{db}
}

func (e *Event) Create(event *models.Event) (err error) {
	_, err = e.db.Index().
		Index("events").
		Type("event").
		BodyJson(e).
		Refresh("true").
		Do(context.Background())

	return
}

func (e *Event) Find(q *models.Query) (events []models.Event, err error) {
	index := e.db.Search().Index("events")

	from, err := strconv.Atoi(q.From)
	if err != nil {
		from = 0
	}

	size, err := strconv.Atoi(q.Size)
	if err != nil {
		size = 10
	}

	if q.UUID != "" {
		cid := elastic.NewTermQuery("cid", q.UUID)
		index.Query(cid)
	}

	if q.Name != "" {
		name := elastic.NewTermQuery("name", q.Name)
		index.Query(name)
	}

	if q.Status != "" {
		status := elastic.NewTermQuery("status", q.Status)
		index.Query(status)
	}

	searchResult, err := index.
		Sort("ocurred_on", true).
		From(from).
		Size(size).
		Do(context.Background())

	if err != nil {
		return
	}

	events, err = parseResult(searchResult)

	return
}

func parseResult(searchResult *elastic.SearchResult) (events []models.Event, err error) {
	for _, hit := range searchResult.Hits.Hits {
		var event models.Event
		json.Unmarshal(*hit.Source, &event)
		events = append(events, event)
	}

	return
}
