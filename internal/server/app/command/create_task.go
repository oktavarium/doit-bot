package command

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/server/app/apperr"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
)

type CreateTask struct {
	OwnerId     string
	ListId      *string
	Name        string
	Description string
}

type createTaskHandler struct {
	domainService planner.DomainService
}

type CreateTaskHandler CommandHandlerHint[CreateTask, string]

func NewCreateTaskHandler(domainService planner.DomainService) CreateTaskHandler {
	return createTaskHandler{
		domainService: domainService,
	}
}

func (h createTaskHandler) Handle(ctx context.Context, cmd CreateTask) (string, error) {
	task, err := h.domainService.NewTask(cmd.OwnerId, cmd.ListId, cmd.Name, cmd.Description)
	if err != nil {
		return "", apperr.FromPlannerError(err)
	}
	if err := h.domainService.SaveTask(ctx, task); err != nil {
		return "", errors.Join(apperr.ErrInternalError, err)
	}
	return task.Id(), nil
}
