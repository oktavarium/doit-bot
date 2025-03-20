package planner

import "errors"

var (
	ErrEmptyTask             = errors.New("empty task")
	ErrInvalidTask           = errors.New("task created outside domain layer")
	ErrTaskNotFound          = errors.New("task not found")
	ErrBadId                 = errors.New("bad id")
	ErrEmptyTaskName         = errors.New("empty task name")
	ErrTooBigTaskName        = errors.New("too big task name")
	ErrTooBigTaskDescription = errors.New("too big task description")
	ErrForbidden             = errors.New("forbidden")
	ErrNothingChaned         = errors.New("nothing changed")
	ErrInternalError         = errors.New("internal error")
	ErrInfrastructureError   = errors.New("infrastructure error")
)
