package tgclient

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/oktavarium/doit-bot/internal/doiterr"
)

func (api *TGClient) SendMessage(
	ctx context.Context,
	messageText string,
	chatID int64,
) error {
	if messageText == "" {
		return doiterr.ErrEmptyParameters
	}

	_, err := api.bot.SendMessage(
		ctx,
		&bot.SendMessageParams{
			ChatID: chatID,
			Text:   messageText,
		},
	)

	if err != nil {
		slog.Error("error sending message", slog.Any("error", err))
		return fmt.Errorf("send message: %w", err)
	}

	return nil
}
