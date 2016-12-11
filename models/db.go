package models

import (
	"gopkg.in/olivere/elastic.v3"
	"log"
)

type DB struct {
	*elastic.Client
}

func NewDB(dbUrl string) (*DB, error) {
	url := elastic.SetURL(dbUrl)
	sniff := elastic.SetSniff(false)
	client, err := elastic.NewClient(sniff, url)
	if err != nil {
		panic(err)
	}

	_, err = client.CreateIndex("events").Do()
	if err != nil {
		log.Print(err)
	}

	return &DB{client}, nil
}
