package datastore

import (
	"net/url"

	"gopkg.in/olivere/elastic.v5"
)

const (
	Elastic = "elastic"
)

func New(dsn string) (*elastic.Client, error) {
	url, err := url.Parse(dsn)
	if err != nil {
		return nil, err
	}

	switch url.Scheme {
	case Elastic:
		c := ESConfig{
			Url:   dsn,
			Sniff: false,
		}

		return NewElastic(c)
	default:
		return nil, ErrUnknownDatabaseProvider
	}
}
