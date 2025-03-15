package query

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/doiterr"
	"github.com/oktavarium/doit-bot/internal/server/app/apperr"
	"github.com/oktavarium/doit-bot/internal/server/domain/users"
)

type GetUserByTgId struct {
	TgId int64
}

type getUserByTgIdHandler struct {
	domainService users.DomainService
}

type GetUserByTgIdHandler QueryHandler[GetUserByTgId, *users.User]

func NewGetUserByTgIdHandler(domainService users.DomainService) GetUserByTgIdHandler {
	return getUserByTgIdHandler{
		domainService: domainService,
	}
}

func (h getUserByTgIdHandler) Handle(ctx context.Context, cmd GetUserByTgId) (*users.User, error) {
	user, err := h.domainService.GetUserByTgId(ctx, cmd.TgId)
	if err != nil {
		if err != nil {
			switch {
			case errors.Is(err, users.ErrBadTgId):
				return nil, doiterr.WrapError(apperr.ErrValidationError, err)
			case errors.Is(err, users.ErrUserNotFound):
				return nil, doiterr.WrapError(apperr.ErrNotFoundError, err)
			default:
				return nil, doiterr.WrapError(apperr.ErrInternalError, err)
			}
		}
	}

	return user, nil
}
