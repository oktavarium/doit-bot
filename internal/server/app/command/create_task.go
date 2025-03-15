package command

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/doiterr"
	"github.com/oktavarium/doit-bot/internal/server/app/apperr"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
)

type CreateTask struct {
	OwnerId     string
	Name        string
	Description string
}

type createTaskHandler struct {
	domainService planner.DomainService
}

type CreateTaskHandler CommandHandlerHint[CreateTask, string]

func NewCreateTaskHandler(domainService planner.DomainService) CreateTaskHandler {
	return createTaskHandler{
		domainService: domainService,
	}
}

func (h createTaskHandler) Handle(ctx context.Context, cmd CreateTask) (string, error) {
	taskId, err := h.domainService.CreateTask(ctx, cmd.OwnerId, cmd.Name, cmd.Description)
	if err != nil {
		switch {
		case errors.Is(err, planner.ErrEmptyTaskName),
			errors.Is(err, planner.ErrTooBigTaskName),
			errors.Is(err, planner.ErrTooBigTaskDescription),
			errors.Is(err, planner.ErrEmptyTaskName):
			return "", doiterr.WrapError(apperr.ErrValidationError, err)
		default:
			return "", doiterr.WrapError(apperr.ErrInternalError, err)
		}
	}
	return taskId, nil
}
