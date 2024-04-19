package id

import (
	"errors"
	"github.com/oklog/ulid/v2"
)

type ID[T any] struct {
	value string
}

// NewID creates an ID with generic type used as a phantom type that
// averts comparing two entities of distinct types.
func NewID[T any]() ID[T] {
	return ID[T]{
		value: ulid.Make().String(),
	}
}

var ErrInvalidULID = errors.New("invalid ULID")

// RestoreID restores an ID.
func RestoreID[T any](value string) (ID[T], error) {
	id, err := ulid.ParseStrict(value)
	if err != nil {
		return ID[T]{}, ErrInvalidULID
	}
	return ID[T]{value: id.String()}, nil
}

func (id ID[T]) Unwrap() string {
	return id.value
}
