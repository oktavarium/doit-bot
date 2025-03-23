package command

import (
	"context"

	"github.com/oktavarium/doit-bot/internal/server/app/apperr"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
)

type DeleteList struct {
	ListId  string
	ActorId string
}

type deleteListHandler struct {
	domainService planner.DomainService
}

type DeleteListHandler CommandHandler[DeleteList]

func NewDeleteListHandler(domainService planner.DomainService) DeleteListHandler {
	return deleteListHandler{
		domainService: domainService,
	}
}

func (h deleteListHandler) Handle(ctx context.Context, cmd DeleteList) error {
	if err := h.domainService.DeleteList(ctx, cmd.ActorId, cmd.ListId); err != nil {
		return apperr.FromPlannerError(err)
	}
	return nil
}
