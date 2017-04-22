package repos

import (
	"context"

	"github.com/rafaeljesus/event-srv/pkg/models"
	client "gopkg.in/olivere/elastic.v5"
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

func (e *Event) Find(query *models.Query) (events []models.Event, err error) {
	index := e.db.Search().Index("events")

	if q.UUID != "" {
		cid := client.NewTermQuery("cid", q.UUID)
		index.Query(cid)
	}

	if q.Name != "" {
		name := client.NewTermQuery("name", q.Name)
		index.Query(name)
	}

	if q.Status != "" {
		status := client.NewTermQuery("status", q.Status)
		index.Query(status)
	}

	searchResult, err := index.
		Sort("ocurred_on", true).
		From(0).
		Size(10).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		return
	}

	events, err = parseResult(searchResult)

	return
}

func parseResult(searchResult *client.SearchResult) (events []Event, err error) {
	for _, hit := range searchResult.Hits.Hits {
		var event Event
		json.Unmarshal(*hit.Source, &event)
		events = append(events, event)
	}

	return
}
