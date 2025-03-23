package query

import (
	"context"

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
		return nil, apperr.FromPlannerError(err)
	}
	return task, nil
}
