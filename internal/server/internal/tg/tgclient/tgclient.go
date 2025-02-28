package tgclient

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
)

type TGClient struct {
	botName string
	bot     *bot.Bot
}

func New(token string) (*TGClient, error) {
	bot, err := bot.New(token)
	if err != nil {
		return nil, fmt.Errorf("create tg bot: %w", err)
	}

	msg, err := bot.GetMe(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error call getMe: %w", err)
	}

	return &TGClient{
		botName: msg.Username,
		bot:     bot,
	}, nil
}
