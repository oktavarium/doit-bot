package handlers

import (
	"context"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handlers) MyChatMember(ctx context.Context, b *bot.Bot, update *models.ChatMemberUpdated) {
	if update.NewChatMember.Member != nil {
		if h.botAddedCallback != nil {
			if err := h.botAddedCallback(ctx, update.Chat.ID, update.From.ID, update.From.Username); err != nil {
				slog.Error("bot added callback", slog.Any("error", err))
			}
		}
	}

	if update.NewChatMember.Left != nil {
		slog.Info(
			"I'm removed from the chat ",
			slog.String("Title:", update.Chat.Title),
			slog.String("BY:", update.From.Username),
		)
	}

	if update.NewChatMember.Banned != nil {
		slog.Info(
			"I'm kicked from the chat ",
			slog.String("Title:", update.Chat.Title),
			slog.String("BY:", update.From.Username),
		)
	}
}
