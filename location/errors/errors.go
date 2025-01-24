package errors

import (
	"errors"
)

// ErrNotFound for location.
var ErrNotFound = errors.New("not found")

// IsNotFound for location.
func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}
