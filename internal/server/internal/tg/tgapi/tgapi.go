package tgapi

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/oktavarium/doit-bot/internal/server/internal/model"
	"github.com/oktavarium/doit-bot/internal/server/internal/tg/tgapi/internal/handlers"
)

type TGAPI struct {
	handlers *handlers.Handlers
	bot      *bot.Bot
}

func New(token string, model *model.Model) (*TGAPI, error) {
	handlers := handlers.New(model)
	opts := []bot.Option{
		bot.WithDefaultHandler(handlers.DefaultHandler),
	}

	bot, err := bot.New(token, opts...)
	if err != nil {
		return nil, fmt.Errorf("create tg bot: %w", err)
	}

	return &TGAPI{
		bot:      bot,
		handlers: handlers,
	}, nil
}

func (api *TGAPI) Serve(ctx context.Context) {
	api.bot.Start(ctx)
}
