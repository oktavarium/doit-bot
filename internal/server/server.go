package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/oktavarium/doit-bot/internal/config"
	"github.com/oktavarium/doit-bot/internal/server/internal/api"
	"github.com/oktavarium/doit-bot/internal/server/internal/bot_api"
	"github.com/oktavarium/doit-bot/internal/server/internal/tg_api"
	"golang.org/x/sync/errgroup"
)

type server struct {
	bot    *bot.Bot
	cfg    *config.Config
	tgAPI  *tg_api.TgAPI
	botAPI *bot_api.BotAPI
	api    *api.API
}

func newServer(cfg *config.Config) (*server, error) {
	tgAPI, err := tg_api.New()
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

	botAPI, err := bot_api.New(context.Background(), bot)
	if err != nil {
		return nil, fmt.Errorf("create bot api: %w", err)
	}

	api := api.New(cfg.GetEndpoint(), cfg.GetToken())

	return &server{
		bot:    bot,
		cfg:    cfg,
		tgAPI:  tgAPI,
		botAPI: botAPI,
		api:    api,
	}, nil
}

func (s *server) serve() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(s.api.Serve)
	go s.bot.Start(ctx)
	if err := eg.Wait(); err != nil {
		return fmt.Errorf("start api: %w", err)
	}

	return nil
}
