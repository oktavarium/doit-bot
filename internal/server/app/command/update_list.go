package command

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/server/app/apperr"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
)

type UpdateList struct {
	ActorId     string
	ListId      string
	Name        *string
	Description *string
}

type updateListHandler struct {
	domainService planner.DomainService
}

type UpdateListHandler CommandHandler[UpdateList]

func NewUpdateListHandler(domainService planner.DomainService) UpdateListHandler {
	return updateListHandler{
		domainService: domainService,
	}
}

func (h updateListHandler) Handle(ctx context.Context, cmd UpdateList) error {
	list, err := h.domainService.GetList(ctx, cmd.ActorId, cmd.ListId)
	if err != nil {
		return apperr.FromPlannerError(err)
	}

	var (
		nameChanged        bool
		descriptionChanged bool
	)

	if cmd.Name != nil {
		if err := list.SetName(cmd.ActorId, *cmd.Name); err != nil {
			if !errors.Is(err, planner.ErrNothingChaned) {
				return apperr.FromPlannerError(err)
			}
		} else {
			nameChanged = true
		}
	}

	if cmd.Description != nil {
		if err := list.SetDescription(cmd.ActorId, *cmd.Description); err != nil {
			if !errors.Is(err, planner.ErrNothingChaned) {
				return apperr.FromPlannerError(err)
			}
		} else {
			descriptionChanged = true
		}
	}

	if nameChanged || descriptionChanged {
		if err := h.domainService.UpdateList(ctx, cmd.ActorId, list); err != nil {
			return apperr.FromPlannerError(err)
		}
	}

	return nil
}
