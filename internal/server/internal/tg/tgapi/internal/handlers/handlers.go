package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/oktavarium/doit-bot/internal/server/internal/model"
)

type Handlers struct {
	model *model.Model
}

func New(model *model.Model) *Handlers {
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
