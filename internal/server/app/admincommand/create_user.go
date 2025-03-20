package admincommand

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/server/app/apperr"
	"github.com/oktavarium/doit-bot/internal/server/domain/users"
)

type CreateUser struct {
	ActorTgId int64
	UserTgId  int64
	ChatTgId  int64
	Username  string
}

type createUserHandler struct {
	adminsMap     map[int64]struct{}
	domainService users.DomainService
}

type CreateUserHandler CommandHandler[CreateUser]

func NewCreateUserHandler(admins []int64, domainService users.DomainService) CreateUserHandler {
	adminsMap := make(map[int64]struct{})
	for _, admin := range admins {
		adminsMap[admin] = struct{}{}
	}
	return createUserHandler{
		domainService: domainService,
		adminsMap:     adminsMap,
	}
}

func (h createUserHandler) Handle(ctx context.Context, cmd CreateUser) error {
	_, ok := h.adminsMap[cmd.ActorTgId]
	if !ok {
		return apperr.ErrForbidden
	}

	if err := h.domainService.CreateUser(ctx, cmd.UserTgId, cmd.ChatTgId, cmd.Username); err != nil {
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
