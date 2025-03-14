package command

import (
	"context"

	"github.com/oktavarium/doit-bot/internal/server/domain/users"
)

type CreateUser struct {
	TgId     int64
	ChatTgId int64
	Username string
}

type createUserHandler struct {
	domainService *users.DomainService
}

type CreateUserHandler CommandHandler[CreateUser]

func NewCreateUserHandler(domainService *users.DomainService) CreateUserHandler {
	return createUserHandler{
		domainService: domainService,
	}
}

func (h createUserHandler) Handle(ctx context.Context, cmd CreateUser) error {
	return h.domainService.CreateUser(ctx, cmd.TgId, cmd.ChatTgId, cmd.Username)
}
