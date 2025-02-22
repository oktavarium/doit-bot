package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/oktavarium/doit-bot/internal/config"
	"github.com/oktavarium/doit-bot/internal/server/internal/bot_api"
	"github.com/oktavarium/doit-bot/internal/server/internal/tg_api"
)

type server struct {
	bot    *bot.Bot
	cfg    *config.Config
	tgAPI  *tg_api.TgAPI
	botAPI *bot_api.BotAPI
}

func newServer(cfg *config.Config) (*server, error) {
	tgAPI, err := tg_api.Init()
	if err != nil {
		return nil, fmt.Errorf("init tg api: %w", err)
	}

	opts := []bot.Option{
		bot.WithDefaultHandler(tgAPI.GetDefaultHandler()),
	}

	bot, err := bot.New(cfg.GetToken(), opts...)
	if err != nil {
		return nil, fmt.Errorf("create tg bot: %w", err)
	}

	botAPI, err := bot_api.New(context.Background(), cfg.GetWebAppURL(), bot)
	if err != nil {
		return nil, fmt.Errorf("create bot api: %w", err)
	}

	return &server{
		bot:    bot,
		cfg:    cfg,
		tgAPI:  tgAPI,
		botAPI: botAPI,
	}, nil
}

func (s *server) serve() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	s.bot.Start(ctx)
}
