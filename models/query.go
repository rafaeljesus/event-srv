package models

type Query struct {
	Cid    string
	Name   string
	Status string
}

func NewQuery(cid, name, status string) *Query {
	return &Query{
		cid,
		name,
		status,
	}
}
