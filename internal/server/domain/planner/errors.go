package planner

import "errors"

var (
	ErrEmptyTask           = errors.New("empty task")
	ErrInvalidTask         = errors.New("task created outside domain layer")
	ErrEmptyList           = errors.New("empty list")
	ErrInvalidList         = errors.New("list created outside domain layer")
	ErrNotFound            = errors.New("not found")
	ErrBadId               = errors.New("bad id")
	ErrEmptyName           = errors.New("empty task name")
	ErrTooBigName          = errors.New("too big task name")
	ErrTooBigDescription   = errors.New("too big task description")
	ErrForbidden           = errors.New("forbidden")
	ErrNothingChaned       = errors.New("nothing changed")
	ErrInternalError       = errors.New("internal error")
	ErrInfrastructureError = errors.New("infrastructure error")
)
