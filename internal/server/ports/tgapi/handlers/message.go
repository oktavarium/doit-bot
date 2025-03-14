package handlers

import (
	"context"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/oktavarium/doit-bot/internal/server/app/command"
)

func (h *Handlers) Message(ctx context.Context, b *bot.Bot, update *models.Message) {
	// пользователь добавил бота и начал с ним взаимодействовать, нажав start
	if update.Text == "/start" && update.From != nil && isBotCommand(update.Entities) {
		if err := h.app.Commands.CreateUser.Handle(
			ctx,
			command.CreateUser{
				TgId:     update.From.ID,
				ChatTgId: update.Chat.ID,
				Username: update.From.Username,
			},
		); err != nil {
			slog.Error("create user", slog.Any("error", err))
		}
	}
}

func isBotCommand(entities []models.MessageEntity) bool {
	for _, ent := range entities {
		if ent.Type == models.MessageEntityTypeBotCommand {
			return true
		}
	}
	return false
}
