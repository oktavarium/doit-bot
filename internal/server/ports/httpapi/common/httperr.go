package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlannerError struct {
	Status int
	Err    error
}

func (p PlannerError) Error() string {
	return p.Err.Error()
}

func NewBadRequestError(err error) PlannerError {
	return PlannerError{
		Status: http.StatusBadRequest,
		Err:    err,
	}
}

func NewUnauthorizedError(err error) PlannerError {
	return PlannerError{
		Status: http.StatusUnauthorized,
		Err:    err,
	}
}

func NewInternalServerError(err error) PlannerError {
	return PlannerError{
		Status: http.StatusInternalServerError,
		Err:    err,
	}
}

func AbortContextWithError(ctx *gin.Context, err PlannerError) {
	_ = ctx.Error(err)
	ctx.Abort()
}

func ErrorToContext(ctx *gin.Context, err PlannerError) {
	_ = ctx.Error(err)
}
