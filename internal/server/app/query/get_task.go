package query

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/doiterr"
	"github.com/oktavarium/doit-bot/internal/server/app/apperr"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
)

type GetTask struct {
	ActorId string
	TaskId  string
}

type getTaskHandler struct {
	domainService planner.DomainService
}

type GetTaskHandler QueryHandler[GetTask, *planner.Task]

func NewGetTaskHandler(domainService planner.DomainService) GetTaskHandler {
	return getTaskHandler{
		domainService: domainService,
	}
}

func (h getTaskHandler) Handle(ctx context.Context, cmd GetTask) (*planner.Task, error) {
	task, err := h.domainService.GetTask(ctx, cmd.ActorId, cmd.TaskId)
	if err != nil {
		switch {
		case errors.Is(err, planner.ErrBadId):
			return nil, doiterr.WrapError(apperr.ErrValidationError, err)
		case errors.Is(err, planner.ErrTaskNotFound):
			return nil, doiterr.WrapError(apperr.ErrNotFoundError, err)
		default:
			return nil, doiterr.WrapError(apperr.ErrInternalError, err)
		}
	}
	return task, nil
}
