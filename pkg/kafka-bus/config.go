package kafkabus

import (
	"time"
)

type Config struct {
	Url      string
	Attempts int
	Timeout  time.Duration
}
