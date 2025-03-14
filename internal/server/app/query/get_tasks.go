package query

import (
	"context"

	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
)

type GetTasks struct {
	UserId string
}

type getTasksHandler struct {
	domainService *planner.DomainService
}

type GetTasksHandler QueryHandler[GetTasks, []*planner.Task]

func NewGetTaskskHandler(domainService *planner.DomainService) GetTasksHandler {
	return getTasksHandler{
		domainService: domainService,
	}
}

func (h getTasksHandler) Handle(ctx context.Context, cmd GetTasks) ([]*planner.Task, error) {
	return h.domainService.GetTasks(ctx, cmd.UserId)
}
