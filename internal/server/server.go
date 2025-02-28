package server

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/oktavarium/doit-bot/internal/config"
	"github.com/oktavarium/doit-bot/internal/server/internal/api"
	"github.com/oktavarium/doit-bot/internal/server/internal/model"
	"github.com/oktavarium/doit-bot/internal/server/internal/storage"
	"github.com/oktavarium/doit-bot/internal/server/internal/tg/tgapi"
	"github.com/oktavarium/doit-bot/internal/server/internal/tg/tgclient"
	"golang.org/x/sync/errgroup"
)

type server struct {
	tgapi *tgapi.TGAPI
	api   *api.API
}

func newServer(cfg *config.Config) (*server, error) {
	tgclient, err := tgclient.New(cfg.GetToken())
	if err != nil {
		return nil, fmt.Errorf("create tg client: %w", err)
	}

	storage, err := storage.New(cfg.GetUri())
	if err != nil {
		return nil, fmt.Errorf("new storage: %w", err)
	}

	model := model.New(tgclient, storage)

	tgapi, err := tgapi.New(cfg.GetToken(), model)
	if err != nil {
		return nil, fmt.Errorf("init tg api: %w", err)
	}

	api := api.New(cfg.GetEndpoint(), cfg.GetToken(), model)

	return &server{
		tgapi: tgapi,
		api:   api,
	}, nil
}

func (s *server) serve() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return s.api.Serve(ctx)
	})
	eg.Go(func() error {
		s.tgapi.Serve(ctx)
		return nil
	})
	if err := eg.Wait(); err != nil {
		return fmt.Errorf("start api: %w", err)
	}

	slog.Info("Goodbye!")
	return nil
}
