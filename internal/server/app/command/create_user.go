package command

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/server/app/apperr"
	"github.com/oktavarium/doit-bot/internal/server/domain/users"
)

type CreateUser struct {
	TgId     int64
	ChatTgId int64
	Username string
}

type createUserHandler struct {
	domainService users.DomainService
}

type CreateUserHandler CommandHandler[CreateUser]

func NewCreateUserHandler(domainService users.DomainService) CreateUserHandler {
	return createUserHandler{
		domainService: domainService,
	}
}

func (h createUserHandler) Handle(ctx context.Context, cmd CreateUser) error {
	if err := h.domainService.CreateUser(ctx, cmd.TgId, cmd.ChatTgId, cmd.Username); err != nil {
		switch {
		case errors.Is(err, users.ErrEmptyUsername),
			errors.Is(err, users.ErrBadTgId):
			return errors.Join(apperr.ErrValidationError, err)
		case errors.Is(err, users.ErrUserExists):
			return errors.Join(apperr.ErrAlreadyExistsError, err)
		default:
			return errors.Join(apperr.ErrInternalError, err)
		}
	}
	return nil
}
