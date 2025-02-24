package bot_api

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
)

type BotAPI struct {
	botName string
	bot     *bot.Bot
}

func New(ctx context.Context, bot *bot.Bot) (*BotAPI, error) {
	msg, err := bot.GetMe(ctx)
	if err != nil {
		return nil, fmt.Errorf("error call getMe: %w", err)
	}

	return &BotAPI{
		botName: msg.Username,
		bot:     bot,
	}, nil
}
