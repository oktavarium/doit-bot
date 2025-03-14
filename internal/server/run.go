package server

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/oktavarium/doit-bot/internal/config"
	"github.com/oktavarium/doit-bot/internal/server/adapters/storage"
	"github.com/oktavarium/doit-bot/internal/server/adapters/tgclient"
	"github.com/oktavarium/doit-bot/internal/server/app"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
	"github.com/oktavarium/doit-bot/internal/server/domain/users"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi"
	"github.com/oktavarium/doit-bot/internal/server/ports/tgapi"
	"golang.org/x/sync/errgroup"
)

// Run - main bot function
func Run() error {
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("get config: %w", err)
	}

	// Init adapters (secondary adapters)
	tgClient, err := tgclient.New(cfg.GetToken())
	if err != nil {
		return fmt.Errorf("create tg client: %w", err)
	}

	storage, err := storage.New(cfg.GetUri())
	if err != nil {
		return fmt.Errorf("new storage: %w", err)
	}

	usersDomainService := users.NewDomainService(storage)
	plannerDomainService := planner.NewDomainService(storage)

	// Init app
	app := app.New(tgClient, plannerDomainService, usersDomainService)

	// Init ports (primary adapters)
	tgAPI, err := tgapi.New(cfg.GetToken(), app)
	if err != nil {
		return fmt.Errorf("init tg api: %w", err)
	}

	httpAPI := httpapi.New(cfg.GetEndpoint(), cfg.GetToken(), app)

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
