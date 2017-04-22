package datastore

import (
	"context"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/olivere/elastic.v5"
)

type ESConfig struct {
	Url   string
	Sniff bool
	Index string
}

func NewElastic(c ESConfig) (conn *elastic.Client, err error) {
	url := elastic.SetURL(c.Url)
	sniff := elastic.SetSniff(c.Sniff)
	conn, err = elastic.NewClient(sniff, url)
	if err != nil {
		return
	}

	_, err = conn.CreateIndex(c.Index).Do(context.Background())
	if err != nil {
		log.WithError(err).Warn("Failed to create index!")
	}

	return
}
