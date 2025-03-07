package tgclient

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/oktavarium/doit-bot/internal/doiterr"
)

func (c *client) SendStartupButton(
	ctx context.Context,
	buttonText string,
	messageText string,
	chatID int64,
) error {
	if buttonText == "" || messageText == "" {
		return doiterr.ErrEmptyParameters
	}

	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{
					Text: buttonText,
					URL:  "https://t.me/" + c.botName + "?startapp=inline",
				},
			},
		},
	}

	_, err := c.bot.SendMessage(
		ctx,
		&bot.SendMessageParams{
			ChatID:      chatID,
			Text:        messageText,
			ReplyMarkup: kb,
		},
	)

	if err != nil {
		slog.Error("error sending message", slog.Any("error", err))
		return fmt.Errorf("send message: %w", err)
	}

	return nil
}
