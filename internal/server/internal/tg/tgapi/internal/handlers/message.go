package handlers

import (
	"context"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handlers) Message(ctx context.Context, b *bot.Bot, update *models.Message) {
	// пользователь добавил бота и начал с ним взаимодействовать, нажав start
	if update.Text == "/start" && update.From != nil && isBotCommand(update.Entities) {
		if err := h.model.CreateUser(
			ctx,
			update.From.ID,
			update.Chat.ID,
			update.From.FirstName,
			update.From.LastName,
			update.From.Username,
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
