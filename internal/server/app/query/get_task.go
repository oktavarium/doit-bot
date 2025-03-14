package query

import (
	"context"

	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
)

type GetTask struct {
	TaskId string
}

type getTaskHandler struct {
	domainService *planner.DomainService
}

type GetTaskHandler QueryHandler[GetTask, *planner.Task]

func NewGetTaskHandler(domainService *planner.DomainService) GetTaskHandler {
	return getTaskHandler{
		domainService: domainService,
	}
}

func (h getTaskHandler) Handle(ctx context.Context, cmd GetTask) (*planner.Task, error) {
	return h.domainService.GetTask(ctx, cmd.TaskId)
}
