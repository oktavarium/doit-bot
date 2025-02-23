package bot_api

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/oktavarium/doit-bot/internal/doiterr"
)

func (api *BotAPI) SendStartupButton(
	ctx context.Context,
	buttonText string,
	messageText string,
	chatID string,
) error {
	if buttonText == "" || messageText == "" || chatID == "" {
		return doiterr.ErrEmptyParameters
	}

	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{
					Text: buttonText,
					URL:  "https://t.me/" + api.botName + "/startapp",
					WebApp: &models.WebAppInfo{
						URL: api.webAppURL,
					}},
			},
		},
	}

	_, err := api.bot.SendMessage(
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
