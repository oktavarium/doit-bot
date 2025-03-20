package apperr

import "errors"

var (
	ErrNotFoundError      = errors.New("not found")
	ErrValidationError    = errors.New("validation error")
	ErrForbidden          = errors.New("forbidden")
	ErrInternalError      = errors.New("internal error")
	ErrAlreadyExistsError = errors.New("already exists")
)
