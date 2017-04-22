package checker

import (
	"context"

	"gopkg.in/olivere/elastic.v5"
)

type elastic struct {
	url string
}

func NewElastic(url string) *elastic {
	return &elastic{url}
}

func (e *elastic) IsAlive() bool {
	url := elastic.SetURL(e.url)
	sniff := elastic.SetSniff(false)
	conn, err = elastic.NewClient(sniff, url)
	if err != nil {
		return false
	}

	defer conn.Stop()

	_, _, err := client.Ping(e.url).Do(context.Background())
	if err != nil {
		return false
	}

	return true
}
