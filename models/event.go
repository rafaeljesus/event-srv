package models

import (
	"encoding/json"
	client "gopkg.in/olivere/elastic.v3"
	"time"
)

type Event struct {
	Cid       int              `json:"cid"`
	Name      string           `json:"name"`
	Status    string           `json:"status"`
	Payload   *json.RawMessage `json:"payload"`
	CreatedAt time.Time        `json:"created_at, omitempty"`
}

func (repo *DB) CreateEvent(e *Event) error {
	_, err := repo.Index().
		Index("events").
		Type("event").
		BodyJson(e).
		Refresh(true).
		Do()

	if err != nil {
		return err
	}

	return nil
}

func (repo *DB) SearchEvents(q *Query, events *[]Event) error {
	index := repo.Search().Index("events")

	if q.Cid != "" {
		cid := client.NewTermQuery("cid", q.Cid)
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
		Sort("timestamp", true).
		From(0).Size(10).
		Pretty(true).
		Do()

	if err != nil {
		return err
	}

	if err := parseResult(searchResult, events); err != nil {
		return err
	}

	return nil
}

func parseResult(searchResult *client.SearchResult, events *[]Event) error {
	for _, hit := range searchResult.Hits.Hits {
		var event Event
		json.Unmarshal(*hit.Source, &event)
		*events = append(*events, event)
	}

	return nil
}
