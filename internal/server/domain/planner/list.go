package planner

import (
	"fmt"
)

type List struct {
	id          string
	ownerId     string
	name        string
	description string
	_valid      bool
}

func newList(
	ownerId string,
	name string,
	description string,
) (*List, error) {
	if err := validateId(ownerId); err != nil {
		return nil, fmt.Errorf("validate owner id: %w", err)
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

	return &List{
		id:          newId,
		ownerId:     ownerId,
		name:        name,
		description: description,
		_valid:      true,
	}, nil
}

func (l *List) Id() string {
	return l.id
}

func (l *List) OwnerId() string {
	return l.ownerId
}

func (l *List) Name() string {
	return l.name
}

func (l *List) Description() string {
	return l.description
}

func (l *List) IsValid() bool {
	return l._valid
}

func (l *List) SetName(actorId string, name string) error {
	if l.ownerId != actorId {
		return ErrForbidden
	}

	if err := validateName(name); err != nil {
		return fmt.Errorf("validate name: %w", err)
	}

	if l.name == name {
		return ErrNothingChaned
	}

	l.name = name
	return nil
}

func (l *List) SetDescription(actorId string, description string) error {
	if err := validateId(actorId); err != nil {
		return fmt.Errorf("validate actor id: %w", err)
	}

	if l.ownerId != actorId {
		return ErrForbidden
	}

	if err := validateDescription(description); err != nil {
		return fmt.Errorf("validate description :%w", err)
	}

	if l.description == description {
		return ErrNothingChaned
	}

	l.description = description
	return nil
}

func RestoreListFromDB(
	id string,
	ownerId string,
	name string,
	description string,
) (*List, error) {
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

	return &List{
		id:          id,
		ownerId:     ownerId,
		name:        name,
		description: description,
		_valid:      true,
	}, nil
}
