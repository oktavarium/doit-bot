package planner

import (
	"errors"

	"github.com/google/uuid"
)

const (
	maxTaskNameLength        = 32
	maxTaskDescriptionLength = 256
)

var (
	ErrEmptyTask             = errors.New("empty task")
	ErrBadId                 = errors.New("bad id")
	ErrEmptyTaskName         = errors.New("empty task name")
	ErrTooBigTaskName        = errors.New("too big task name")
	ErrTooBigTaskDescription = errors.New("too big task description")
	ErrForbidden             = errors.New("forbidden")
	ErrNothingChaned         = errors.New("nothing changed")
)

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
		// TODO: return error in production
		panic(err)
	}

	return nil
}

func validateOwnerId(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		// TODO: return error in production
		panic(err)
	}

	return nil
}
