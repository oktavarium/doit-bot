package handlers

import (
	"context"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func InlineQuery(ctx context.Context, b *bot.Bot, update *models.Update) {
	slog.Info("INILNE QUERY")
}
