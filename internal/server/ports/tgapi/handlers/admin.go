package handlers

import (
	"context"
	"log/slog"
	"strconv"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/oktavarium/doit-bot/internal/server/app/admincommand"
	"github.com/oktavarium/doit-bot/internal/server/app/adminquery"
)

func (h *Handlers) AdminHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	isAdmin, err := h.app.AdminQueries.IsAdmin.Handle(ctx, adminquery.IsAdmin{TgId: update.Message.From.ID})
	if err != nil || !isAdmin {
		return
	}

	args := strings.TrimSpace(strings.TrimPrefix(update.Message.Text, "/admin"))
	argsSlice := strings.Split(args, " ")
	if len(argsSlice) != 2 {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Not enough args",
		})
		return
	}

	cmd := argsSlice[0]
	id, err := strconv.ParseInt(argsSlice[1], 10, 64)
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Wrong args",
		})
		return
	}
	switch cmd {
	case "create":
		if err := h.app.AdminCommands.CreateUser.Handle(
			ctx,
			admincommand.CreateUser{
				ActorTgId: update.Message.From.ID,
				UserTgId:  id,
				ChatTgId:  update.Message.Chat.ID,
				Username:  update.Message.From.Username,
			},
		); err != nil {
			slog.Error("create user", slog.Any("error", err))
		}
	default:
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Unknown command",
		})
		return
	}

}
