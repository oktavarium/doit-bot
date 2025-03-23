package command

import (
	"context"

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
		return apperr.FromPlannerError(err)
	}
	return nil
}
