package handlers

import (
	"context"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func CallbackQuery(ctx context.Context, b *bot.Bot, update *models.Update) {
	slog.Info("CALLBACK QUERY")
}
