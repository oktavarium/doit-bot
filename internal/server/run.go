package server

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/oktavarium/doit-bot/internal/config"
	"github.com/oktavarium/doit-bot/internal/server/adapters/http_api"
	"github.com/oktavarium/doit-bot/internal/server/adapters/storage"
	"github.com/oktavarium/doit-bot/internal/server/adapters/tg_api"
	"github.com/oktavarium/doit-bot/internal/server/adapters/tg_client"
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
	tgclient, err := tg_client.New(cfg.GetToken())
	if err != nil {
		return fmt.Errorf("create tg client: %w", err)
	}

	storage, err := storage.New(cfg.GetUri())
	if err != nil {
		return fmt.Errorf("new storage: %w", err)
	}

	// Init domain model
	model := model.New(tgclient, storage)

	// Init incoming adapters
	tg_server, err := tg_api.New(cfg.GetToken(), model)
	if err != nil {
		return fmt.Errorf("init tg api: %w", err)
	}

	http_server := http_api.New(cfg.GetEndpoint(), cfg.GetToken(), model)

	// Start server
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return http_server.Serve(ctx)
	})
	eg.Go(func() error {
		return tg_server.Serve(ctx)
	})
	if err := eg.Wait(); err != nil {
		return fmt.Errorf("start api: %w", err)
	}

	slog.Info("Goodbye!")
	return nil
}
