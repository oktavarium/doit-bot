package handlers

import (
	"context"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handlers) MyChatMember(ctx context.Context, b *bot.Bot, update *models.ChatMemberUpdated) {
	if update.NewChatMember.Member != nil {
		if update.Chat.Type == models.ChatTypeGroup {
			slog.Info(
				"I'm added to the chat ",
				slog.String("Title:", update.Chat.Title),
				slog.String("BY:", update.From.Username),
			)
			// if err := h.model.CreateGroup(
			// 	ctx,
			// 	update.From.ID,
			// 	update.Chat.ID,
			// 	update.Chat.Title,
			// ); err != nil {
			// 	slog.Error("create group", slog.Any("error", err))
			// 	return
			// }
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
