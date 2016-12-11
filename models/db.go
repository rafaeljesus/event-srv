package models

import (
	"context"
	log "github.com/Sirupsen/logrus"
	"gopkg.in/olivere/elastic.v5"
)

type DB struct {
	*elastic.Client
}

func NewDB(dbUrl string) (*DB, error) {
	url := elastic.SetURL(dbUrl)
	sniff := elastic.SetSniff(false)
	conn, err := elastic.NewClient(sniff, url)
	if err != nil {
		return nil, err
	}

	_, err = conn.CreateIndex("events").Do(context.Background())
	if err != nil {
		log.WithError(err).Error("Failed to create index!")
	}

	return &DB{conn}, nil
}
