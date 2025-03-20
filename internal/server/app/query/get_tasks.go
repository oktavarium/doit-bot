package query

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/server/app/apperr"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
)

type GetTasks struct {
	UserId string
}

type getTasksHandler struct {
	domainService planner.DomainService
}

type GetTasksHandler QueryHandler[GetTasks, []*planner.Task]

func NewGetTaskskHandler(domainService planner.DomainService) GetTasksHandler {
	return getTasksHandler{
		domainService: domainService,
	}
}

func (h getTasksHandler) Handle(ctx context.Context, cmd GetTasks) ([]*planner.Task, error) {
	tasks, err := h.domainService.GetTasks(ctx, cmd.UserId)
	if err != nil {
		if err != nil {
			switch {
			case errors.Is(err, planner.ErrBadId):
				return nil, errors.Join(apperr.ErrValidationError, err)
			default:
				return nil, errors.Join(apperr.ErrInternalError, err)
			}
		}
	}

	return tasks, nil
}
