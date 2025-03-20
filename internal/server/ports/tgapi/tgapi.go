package tgapi

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/oktavarium/doit-bot/internal/server/app"
	"github.com/oktavarium/doit-bot/internal/server/ports/tgapi/handlers"
)

type TGAPI struct {
	handlers *handlers.Handlers
	bot      *bot.Bot
}

func New(token string, app *app.App) (*TGAPI, error) {
	handlers := handlers.New(app)
	opts := []bot.Option{
		bot.WithDefaultHandler(handlers.DefaultHandler),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		return nil, fmt.Errorf("create tg bot: %w", err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handlers.StartHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/admin", bot.MatchTypePrefix, handlers.AdminHandler)

	return &TGAPI{
		bot:      b,
		handlers: handlers,
	}, nil
}

func (api *TGAPI) Serve(ctx context.Context) error {
	api.bot.Start(ctx)
	return nil
}
