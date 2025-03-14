package httpapi

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/app"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/handlers"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/middleware"
)

type API struct {
	router   *gin.Engine
	endpoint string
	handlers *handlers.Handlers
}

func New(endpoint string, token string, app *app.App) *API {
	router := gin.Default()
	router.ContextWithFallback = true

	middleware.Init(router, token, app)
	handlers := handlers.New(router, token, app)

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
