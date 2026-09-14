package config

import (
	"fmt"
	"time"
)

type ErrInvalidConfiguration struct {
	Key    string
	Value  any
	Reason string
	When   time.Time
}

func (e *ErrInvalidConfiguration) Error() string {
	return fmt.Sprintf("Failed to initialize application: environment variable %s (value: [%v]) is invalid: %s", e.Key, e.Value, e.Reason)
}
