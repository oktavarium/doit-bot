package planner

import (
	"fmt"
)

type Task struct {
	id          string
	ownerId     string
	listId      string
	name        string
	description string
	status      bool
	_valid      bool
}

func newTask(
	ownerId string,
	listId *string,
	name string,
	description string,
) (*Task, error) {
	if err := validateId(ownerId); err != nil {
		return nil, fmt.Errorf("validate owner id: %w", err)
	}

	var list string
	if listId != nil {
		if err := validateId(*listId); err != nil {
			return nil, fmt.Errorf("validate list id: %w", err)
		}
		list = *listId
	}

	if err := validateName(name); err != nil {
		return nil, fmt.Errorf("validate name: %w", err)
	}

	if err := validateDescription(description); err != nil {
		return nil, fmt.Errorf("validate description: %w", err)
	}

	newId, err := generateId()
	if err != nil {
		return nil, fmt.Errorf("generate id: %w", err)
	}

	return &Task{
		id:          newId,
		ownerId:     ownerId,
		listId:      list,
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

func (t *Task) ListId() string {
	return t.listId
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
	if err := validateId(actorId); err != nil {
		return fmt.Errorf("validate actor id: %w", err)
	}

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
	if err := validateId(actorId); err != nil {
		return fmt.Errorf("validate actor id: %w", err)
	}

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
	if err := validateId(actorId); err != nil {
		return fmt.Errorf("validate actor id: %w", err)
	}

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

func (t *Task) SetListId(actorId string, listId string) error {
	if err := validateId(actorId); err != nil {
		return fmt.Errorf("validate actor id: %w", err)
	}

	if err := validateId(listId); err != nil {
		return fmt.Errorf("validate list id: %w", err)
	}

	if t.ownerId != actorId {
		return ErrForbidden
	}

	if t.listId == listId {
		return ErrNothingChaned
	}

	t.listId = listId
	return nil
}

func RestoreTaskFromDB(
	id string,
	ownerId string,
	listId string,
	name string,
	description string,
	status bool,
) (*Task, error) {
	if err := validateId(id); err != nil {
		return nil, fmt.Errorf("validate task id: %w", err)
	}

	if err := validateId(ownerId); err != nil {
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
		listId:      listId,
		name:        name,
		description: description,
		status:      status,
		_valid:      true,
	}, nil
}
