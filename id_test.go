package id

import (
	"errors"
	"testing"
)

type User struct{}

func TestRestoreID(t *testing.T) {
	oid := NewID[User]()
	rid, err := RestoreID[User](oid.Unwrap())
	if err != nil {
		t.Errorf("RestoreID failed with error: %v", err)
	}
	if oid != rid {
		t.Error("Restored ID does not match the original ID")
	}
}

func TestRestoreInvalidID(t *testing.T) {
	_, err := RestoreID[User]("invalid-ulid")
	if !errors.Is(err, ErrInvalidULID) {
		t.Error("RestoreID should have returned an error for an invalid ULID")
	}
}
