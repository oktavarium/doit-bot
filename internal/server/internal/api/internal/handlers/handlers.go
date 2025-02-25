package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/internal/api/internal/middleware/auth"
)

type Handlers struct {
	router *gin.Engine
	token  string
}

func New(router *gin.Engine, token string) *Handlers {
	h := &Handlers{
		router: router,
		token:  token,
	}

	h.init()

	return h
}

func (h *Handlers) init() {
	apiGroup := h.router.Group("/api", auth.Middleware(h.token))
	apiGroup.POST("/", h.CreateData)
	apiGroup.GET("/", h.GetData)

}
