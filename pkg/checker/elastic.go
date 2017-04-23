package checker

import (
	"context"

	client "gopkg.in/olivere/elastic.v5"
)

type elastic struct {
	url string
}

func NewElastic(url string) *elastic {
	return &elastic{url}
}

func (e *elastic) IsAlive() bool {
	url := client.SetURL(e.url)
	sniff := client.SetSniff(false)
	conn, err := client.NewClient(sniff, url)
	if err != nil {
		return false
	}

	defer conn.Stop()

	_, _, err = conn.Ping(e.url).Do(context.Background())
	if err != nil {
		return false
	}

	return true
}
