package common

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/app/apperr"
)

type Error struct {
	Status int
	Err    error
}

func (p Error) Error() string {
	return p.Err.Error()
}

func FromAppError(err error) Error {
	switch {
	case errors.Is(err, apperr.ErrNotFoundError):
		return NewNotFoundError(err)
	case errors.Is(err, apperr.ErrValidationError):
		return NewBadRequestError(err)
	case errors.Is(err, apperr.ErrForbidden):
		return NewForbiddenError(err)
	case errors.Is(err, apperr.ErrAlreadyExistsError):
		return NewConflictError(err)
	default:
		return NewInternalServerError(err)
	}
}

func NewBadRequestError(err error) Error {
	return Error{
		Status: http.StatusBadRequest,
		Err:    err,
	}
}

func NewUnauthorizedError(err error) Error {
	return Error{
		Status: http.StatusUnauthorized,
		Err:    err,
	}
}

func NewForbiddenError(err error) Error {
	return Error{
		Status: http.StatusForbidden,
		Err:    err,
	}
}

func NewInternalServerError(err error) Error {
	return Error{
		Status: http.StatusInternalServerError,
		Err:    err,
	}
}

func NewNotFoundError(err error) Error {
	return Error{
		Status: http.StatusNotFound,
		Err:    err,
	}
}

func NewConflictError(err error) Error {
	return Error{
		Status: http.StatusConflict,
		Err:    err,
	}
}

func AbortContextWithError(ctx *gin.Context, err Error) {
	_ = ctx.Error(err)
	ctx.Abort()
}

func ErrorToContext(ctx *gin.Context, err Error) {
	_ = ctx.Error(err)
}
