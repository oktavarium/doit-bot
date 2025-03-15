package planner

import (
	"github.com/google/uuid"
	"github.com/oktavarium/doit-bot/internal/doiterr"
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
	if id == "" {
		return nil
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return doiterr.WrapError(ErrInternalError, err)
	}

	return nil
}

func validateOwnerId(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return doiterr.WrapError(ErrInternalError, err)
	}

	return nil
}
