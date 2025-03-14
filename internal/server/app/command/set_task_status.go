package command

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
)

type SetTaskStatus struct {
	TaskId string
	Status bool
}

type setTaskStatusHandler struct {
	domainService *planner.DomainService
}

type SetTaskStatusHandler CommandHandler[SetTaskStatus]

func NewSetTaskStatusHandler(domainService *planner.DomainService) SetTaskStatusHandler {
	return setTaskStatusHandler{
		domainService: domainService,
	}
}

func (h setTaskStatusHandler) Handle(ctx context.Context, cmd SetTaskStatus) error {
	task, err := h.domainService.GetTask(ctx, cmd.TaskId)
	if err != nil {
		return fmt.Errorf("create task: %w", err)
	}

	if err := task.SetStatus(cmd.Status); err != nil {
		return fmt.Errorf("set status: %w", err)
	}

	return h.domainService.UpdateTask(ctx, task)
}
