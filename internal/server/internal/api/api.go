package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/internal/api/internal/handlers"
	"github.com/oktavarium/doit-bot/internal/server/internal/api/internal/middleware"
)

type API struct {
	router   *gin.Engine
	endpoint string
	handlers *handlers.Handlers
}

func New(endpoint string, token string) *API {
	router := gin.Default()
	router.ContextWithFallback = true

	middleware.Init(router, token)
	handlers := handlers.New(router)

	return &API{
		router:   router,
		endpoint: endpoint,
		handlers: handlers,
	}

}

func (api *API) Serve() error {
	server := &http.Server{
		Addr:    api.endpoint,
		Handler: api.router,
	}

	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("listen and serve: %w,", err)
	}

	return nil
}
