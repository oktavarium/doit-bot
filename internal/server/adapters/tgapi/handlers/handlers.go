package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/oktavarium/doit-bot/internal/server/ports"
)

type Handlers struct {
	model ports.Model
}

func New(model ports.Model) *Handlers {
	return &Handlers{
		model: model,
	}
}

func (h *Handlers) DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	switch {
	case update.Message != nil:
		h.Message(ctx, b, update.Message)
	case update.MyChatMember != nil:
		h.MyChatMember(ctx, b, update.MyChatMember)
	case update.ChatMember != nil:
		ChatMember(ctx, b, update)
	}
}
