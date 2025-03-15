package command

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/doiterr"
	"github.com/oktavarium/doit-bot/internal/server/app/apperr"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
)

type DeleteTask struct {
	TaskId  string
	ActorId string
}

type deleteTaskHandler struct {
	domainService planner.DomainService
}

type DeleteTaskHandler CommandHandler[DeleteTask]

func NewDeleteTaskHandler(domainService planner.DomainService) DeleteTaskHandler {
	return deleteTaskHandler{
		domainService: domainService,
	}
}

func (h deleteTaskHandler) Handle(ctx context.Context, cmd DeleteTask) error {
	if err := h.domainService.DeleteTask(ctx, cmd.ActorId, cmd.TaskId); err != nil {
		switch {
		case errors.Is(err, planner.ErrBadId):
			return doiterr.WrapError(apperr.ErrValidationError, err)
		case errors.Is(err, planner.ErrTaskNotFound):
			return doiterr.WrapError(apperr.ErrNotFoundError, err)
		default:
			return doiterr.WrapError(apperr.ErrInternalError, err)
		}
	}
	return nil
}
