package planner

import (
	"errors"

	"github.com/google/uuid"
)

const (
	maxTaskNameLength        = 64
	maxTaskDescriptionLength = 1024
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

func validateName(name string) error {
	if name == "" {
		return ErrEmptyTaskName
	}

	if len(name) > maxTaskNameLength {
		return ErrTooBigTaskName
	}

	return nil
}

func validateDescription(description string) error {
	if len(description) > maxTaskDescriptionLength {
		return ErrTooBigTaskDescription
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

func validateOwnerId(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return errors.Join(ErrInternalError, err)
	}

	return nil
}
