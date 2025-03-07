package server

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/oktavarium/doit-bot/internal/config"
	"github.com/oktavarium/doit-bot/internal/server/adapters/httpapi"
	"github.com/oktavarium/doit-bot/internal/server/adapters/storage"
	"github.com/oktavarium/doit-bot/internal/server/adapters/tgapi"
	"github.com/oktavarium/doit-bot/internal/server/adapters/tgclient"
	"github.com/oktavarium/doit-bot/internal/server/model"
	"golang.org/x/sync/errgroup"
)

// Run - main bot function
func Run() error {
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("get config: %w", err)
	}

	// Init outcoming adapters
	tgClient, err := tgclient.New(cfg.GetToken())
	if err != nil {
		return fmt.Errorf("create tg client: %w", err)
	}

	storage, err := storage.New(cfg.GetUri())
	if err != nil {
		return fmt.Errorf("new storage: %w", err)
	}

	// Init domain model
	model := model.New(tgClient, storage)

	// Init incoming adapters
	tgAPI, err := tgapi.New(cfg.GetToken(), model)
	if err != nil {
		return fmt.Errorf("init tg api: %w", err)
	}

	httpAPI := httpapi.New(cfg.GetEndpoint(), cfg.GetToken(), model)

	// Start server
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return httpAPI.Serve(ctx)
	})
	eg.Go(func() error {
		return tgAPI.Serve(ctx)
	})
	if err := eg.Wait(); err != nil {
		return fmt.Errorf("start api: %w", err)
	}

	slog.Info("Goodbye!")
	return nil
}
