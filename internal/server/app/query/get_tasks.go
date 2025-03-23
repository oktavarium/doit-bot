package query

import (
	"context"

	"github.com/oktavarium/doit-bot/internal/server/app/apperr"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
)

type GetTasks struct {
	UserId string
	ListId *string
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
	if cmd.ListId == nil {
		tasks, err := h.domainService.GetTasks(ctx, cmd.UserId)
		if err != nil {
			return nil, apperr.FromPlannerError(err)
		}

		return tasks, nil
	}

	tasks, err := h.domainService.GetListTasks(ctx, cmd.UserId, *cmd.ListId)
	if err != nil {
		return nil, apperr.FromPlannerError(err)
	}

	return tasks, nil

}
