package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/internal/api/internal/middleware/auth"
	"github.com/oktavarium/doit-bot/internal/server/internal/storage"
)

type Handlers struct {
	router  *gin.Engine
	token   string
	storage storage.Storage
}

func New(router *gin.Engine, token string, storage storage.Storage) *Handlers {
	h := &Handlers{
		router:  router,
		token:   token,
		storage: storage,
	}

	h.init()

	return h
}

func (h *Handlers) init() {
	apiGroup := h.router.Group("/api", auth.Middleware(h.token))
	apiGroup.POST("/", h.CreateData)
	apiGroup.GET("/", h.GetData)

}
