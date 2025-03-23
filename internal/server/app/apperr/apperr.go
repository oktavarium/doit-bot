package apperr

import (
	"errors"

	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
	"github.com/oktavarium/doit-bot/internal/server/domain/users"
)

var (
	ErrNotFoundError      = errors.New("not found")
	ErrValidationError    = errors.New("validation error")
	ErrForbidden          = errors.New("forbidden")
	ErrInternalError      = errors.New("internal error")
	ErrAlreadyExistsError = errors.New("already exists")
)

func FromPlannerError(err error) error {
	switch {
	case errors.Is(err, planner.ErrEmptyName),
		errors.Is(err, planner.ErrTooBigName),
		errors.Is(err, planner.ErrBadId),
		errors.Is(err, planner.ErrTooBigDescription):
		return errors.Join(ErrValidationError, err)
	case errors.Is(err, planner.ErrNotFound):
		return errors.Join(ErrNotFoundError, err)
	case errors.Is(err, planner.ErrForbidden):
		return errors.Join(ErrForbidden, err)
	default:
		return errors.Join(ErrInternalError, err)
	}
}

func FromUsersError(err error) error {
	switch {
	case errors.Is(err, users.ErrEmptyUsername),
		errors.Is(err, users.ErrBadTgId):
		return errors.Join(ErrValidationError, err)
	case errors.Is(err, users.ErrUserExists):
		return errors.Join(ErrAlreadyExistsError, err)
	case errors.Is(err, users.ErrUserNotFound):
		return errors.Join(ErrNotFoundError, err)
	default:
		return errors.Join(ErrInternalError, err)
	}
}
