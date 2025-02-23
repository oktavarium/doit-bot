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
	apiGroup.POST("/api", h.CreateData)
	apiGroup.GET("/api", h.GetData)

	h.router.LoadHTMLGlob("../web/templates/*")
	h.router.GET("/", h.Main)
	h.router.Static("/static", "../web/static")
}
