package greetings_sdk

import (
	"fmt"
	"time"
)

// ErrInvalidStatus - ошибка при неожиданном HTTP-статусе
type ErrInvalidStatus struct {
	StatusCode int
	Body       string
}

func (e *ErrInvalidStatus) Error() string {
	return fmt.Sprintf("unexpected status %d: %s", e.StatusCode, e.Body)
}

// ErrTimeout - ошибка таймаута
type ErrTimeout struct {
	Timeout time.Duration
}

func (e *ErrTimeout) Error() string {
	return fmt.Sprintf("request timeout after %v", e.Timeout)
}
