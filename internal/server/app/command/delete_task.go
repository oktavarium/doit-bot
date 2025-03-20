package command

import (
	"context"
	"errors"

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
			return errors.Join(apperr.ErrValidationError, err)
		case errors.Is(err, planner.ErrTaskNotFound):
			return errors.Join(apperr.ErrNotFoundError, err)
		default:
			return errors.Join(apperr.ErrInternalError, err)
		}
	}
	return nil
}
