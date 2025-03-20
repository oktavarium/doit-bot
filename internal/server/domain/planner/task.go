package planner

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Task struct {
	id          string
	ownerId     string
	name        string
	description string
	status      bool
	_valid      bool
}

func newTask(
	ownerId string,
	name string,
	description string,
) (*Task, error) {
	if err := validateOwnerId(ownerId); err != nil {
		return nil, fmt.Errorf("validate owner id: %w", err)
	}

	if err := validateName(name); err != nil {
		return nil, fmt.Errorf("validate name: %w", err)
	}

	if err := validateDescription(name); err != nil {
		return nil, fmt.Errorf("validate description: %w", err)
	}

	newId, err := generateId()
	if err != nil {
		return nil, fmt.Errorf("generate id: %w", err)
	}

	return &Task{
		id:          newId,
		ownerId:     ownerId,
		name:        name,
		description: description,
		_valid:      true,
	}, nil
}

func (t *Task) Id() string {
	return t.id
}

func (t *Task) OwnerId() string {
	return t.ownerId
}

func (t *Task) Name() string {
	return t.name
}

func (t *Task) Description() string {
	return t.description
}

func (t *Task) Status() bool {
	return t.status
}

func (t *Task) IsValid() bool {
	return t._valid
}

func (t *Task) SetStatus(actorId string, status bool) error {
	if t.ownerId != actorId {
		return ErrForbidden
	}

	if t.status == status {
		return ErrNothingChaned
	}

	t.status = status
	return nil
}

func (t *Task) SetName(actorId string, name string) error {
	if t.ownerId != actorId {
		return ErrForbidden
	}

	if err := validateName(name); err != nil {
		return fmt.Errorf("validate name: %w", err)
	}

	if t.name == name {
		return ErrNothingChaned
	}

	t.name = name
	return nil
}

func (t *Task) SetDescription(actorId string, description string) error {
	if t.ownerId != actorId {
		return ErrForbidden
	}

	if err := validateDescription(description); err != nil {
		return fmt.Errorf("validate description :%w", err)
	}

	if t.description == description {
		return ErrNothingChaned
	}

	t.description = description
	return nil
}

func generateId() (string, error) {
	newId, err := uuid.NewV7()
	if err != nil {
		return "", errors.Join(ErrInternalError, err)
	}

	return newId.String(), nil
}

func RestoreTaskFromDB(
	id string,
	ownerId string,
	name string,
	description string,
	status bool,
) (*Task, error) {
	if err := validateId(id); err != nil {
		return nil, fmt.Errorf("validate task id: %w", err)
	}

	if err := validateOwnerId(ownerId); err != nil {
		return nil, fmt.Errorf("validate owner id: %w", err)
	}

	if err := validateName(name); err != nil {
		return nil, fmt.Errorf("validate name: %w", err)
	}

	if err := validateDescription(name); err != nil {
		return nil, fmt.Errorf("validate description: %w", err)
	}

	return &Task{
		id:          id,
		ownerId:     ownerId,
		name:        name,
		description: description,
		status:      status,
		_valid:      true,
	}, nil
}
