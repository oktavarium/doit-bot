package tgclient

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/oktavarium/doit-bot/internal/doiterr"
)

func (c *client) SendMessage(
	ctx context.Context,
	messageText string,
	chatID int64,
) error {
	if messageText == "" {
		return doiterr.ErrEmptyParameters
	}

	_, err := c.bot.SendMessage(
		ctx,
		&bot.SendMessageParams{
			ChatID: chatID,
			Text:   messageText,
		},
	)

	if err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	return nil
}
