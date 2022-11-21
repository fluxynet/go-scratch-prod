package go_scratch_prod

import (
	"errors"
)

var (
	// ErrNotImplemented indicates a missing feature from the application
	ErrNotImplemented = errors.New("feature not implemented")

	// ErrEntityNotFound indicates a specific entity could not be found
	ErrEntityNotFound = errors.New("entity not found")

	// ErrInvalidRequest means a request is either nil or not appropriate for the requested action
	ErrInvalidRequest = errors.New("request is invalid")
)
