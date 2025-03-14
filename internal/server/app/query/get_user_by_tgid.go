package query

import (
	"context"

	"github.com/oktavarium/doit-bot/internal/server/domain/users"
)

type GetUserByTgId struct {
	TgId int64
}

type getUserByTgIdHandler struct {
	domainService *users.DomainService
}

type GetUserByTgIdHandler QueryHandler[GetUserByTgId, *users.User]

func NewGetUserByTgIdHandler(domainService *users.DomainService) GetUserByTgIdHandler {
	return getUserByTgIdHandler{
		domainService: domainService,
	}
}

func (h getUserByTgIdHandler) Handle(ctx context.Context, cmd GetUserByTgId) (*users.User, error) {
	return h.domainService.GetUserByTgId(ctx, cmd.TgId)
}
