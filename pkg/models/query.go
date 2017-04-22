package models

type Query struct {
	UUID   string
	Name   string
	Status string
}

func NewQuery(uuid, name, status string) *Query {
	return &Query{uuid, name, status}
}
