package command

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/server/app/apperr"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
)

type CreateList struct {
	OwnerId     string
	Name        string
	Description string
}

type createListHandler struct {
	domainService planner.DomainService
}

type CreateListHandler CommandHandlerHint[CreateList, string]

func NewCreateListHandler(domainService planner.DomainService) CreateListHandler {
	return createListHandler{
		domainService: domainService,
	}
}

func (h createListHandler) Handle(ctx context.Context, cmd CreateList) (string, error) {
	task, err := h.domainService.NewList(cmd.OwnerId, cmd.Name, cmd.Description)
	if err != nil {
		return "", apperr.FromPlannerError(err)
	}
	if err := h.domainService.SaveList(ctx, task); err != nil {
		return "", errors.Join(apperr.ErrInternalError, err)
	}
	return task.Id(), nil
}
