package internal

import "errors"

func newError(msg string) error {
	return errors.New(msg)
}

var (
	ErrNameUnique = newError("Name must be unique.")
)
