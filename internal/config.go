package internal

import (
	"os"
	"time"
)

// Config .
type Config struct {
	Input       *os.File
	Output      *os.File
	Concurrency int
	RateLimit   int
	Timeout     time.Duration
}
