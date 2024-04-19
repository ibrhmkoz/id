package id

import (
	"errors"
	"testing"
)

type User struct{}

func TestRestoreID(t *testing.T) {
	oid := New[User]()
	rid, err := Restore[User](oid.Unwrap())
	if err != nil {
		t.Errorf("Restore failed with error: %v", err)
	}
	if oid != rid {
		t.Error("Restored ID does not match the original ID")
	}
}

func TestRestoreInvalidID(t *testing.T) {
	_, err := Restore[User]("invalid-ulid")
	if !errors.Is(err, ErrInvalidULID) {
		t.Error("Restore should have returned an error for an invalid ULID")
	}
}
