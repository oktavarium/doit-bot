package httpapi

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/adapters/httpapi/handlers"
	"github.com/oktavarium/doit-bot/internal/server/adapters/httpapi/middleware"
	"github.com/oktavarium/doit-bot/internal/server/ports"
)

type API struct {
	router   *gin.Engine
	endpoint string
	handlers *handlers.Handlers
}

func New(endpoint string, token string, model ports.Model) *API {
	router := gin.Default()
	router.ContextWithFallback = true

	middleware.Init(router, token)
	handlers := handlers.New(router, token, model)

	return &API{
		router:   router,
		endpoint: endpoint,
		handlers: handlers,
	}

}

func (api *API) Serve(ctx context.Context) error {
	server := &http.Server{
		Addr:    api.endpoint,
		Handler: api.router,
	}

	go func(ctx context.Context) {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			slog.Error("shutdown server", slog.Any("error", err))
		}
	}(ctx)

	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("listen and serve: %w,", err)
	}

	return nil
}
