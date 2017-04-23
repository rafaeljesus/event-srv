package models

import (
	"encoding/json"
	"time"
)

type Event struct {
	UUID      string           `json:"uuid"`
	Name      string           `json:"name"`
	Status    string           `json:"status"`
	Payload   *json.RawMessage `json:"payload"`
	OcurredOn time.Time        `json:"ocurred_on,omitempty"`
}
