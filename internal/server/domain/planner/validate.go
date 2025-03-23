package planner

import (
	"errors"

	"github.com/google/uuid"
)

const (
	maxTaskNameLength        = 64
	maxTaskDescriptionLength = 256
)

func isTaskValid(t *Task) error {
	if t == nil {
		return ErrEmptyTask
	}

	if !t.IsValid() {
		return ErrInvalidTask
	}

	return nil
}

func isListValid(l *List) error {
	if l == nil {
		return ErrEmptyList
	}

	if !l.IsValid() {
		return ErrInvalidList
	}

	return nil
}

func validateName(name string) error {
	if name == "" {
		return ErrEmptyName
	}

	if len(name) > maxTaskNameLength {
		return ErrTooBigName
	}

	return nil
}

func validateDescription(description string) error {
	if len(description) > maxTaskDescriptionLength {
		return ErrTooBigDescription
	}
	return nil
}

func validateId(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return errors.Join(ErrBadId, err)
	}

	return nil
}
