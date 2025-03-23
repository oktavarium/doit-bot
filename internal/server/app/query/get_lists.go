package query

import (
	"context"

	"github.com/oktavarium/doit-bot/internal/server/app/apperr"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
)

type GetLists struct {
	UserId string
}

type getListsHandler struct {
	domainService planner.DomainService
}

type GetListsHandler QueryHandler[GetLists, []*planner.List]

func NewGetListskHandler(domainService planner.DomainService) GetListsHandler {
	return getListsHandler{
		domainService: domainService,
	}
}

func (h getListsHandler) Handle(ctx context.Context, cmd GetLists) ([]*planner.List, error) {
	tasks, err := h.domainService.GetLists(ctx, cmd.UserId)
	if err != nil {
		return nil, apperr.FromPlannerError(err)
	}

	return tasks, nil
}
