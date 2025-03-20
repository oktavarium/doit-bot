package handlers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/oktavarium/doit-bot/internal/server/app/apperr"
	"github.com/oktavarium/doit-bot/internal/server/app/command"
)

func (h *Handlers) StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if err := h.app.Commands.CreateUser.Handle(
		ctx,
		command.CreateUser{
			TgId:     update.Message.From.ID,
			ChatTgId: update.Message.Chat.ID,
			Username: update.Message.From.Username,
		},
	); err != nil {
		if errors.Is(err, apperr.ErrAlreadyExistsError) {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   fmt.Sprintf("You are already registered. TgId: %d", update.Message.From.ID),
			})
			return
		}

		slog.Error("create user", slog.Any("error", err))
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Error creating user",
		})
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Now you can use miniapp. TgId: %d", update.Message.From.ID),
	})
}
