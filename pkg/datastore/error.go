package datastore

import "errors"

var (
	ErrUnknownDatabaseProvider = errors.New("Unknown database provider type")
)
