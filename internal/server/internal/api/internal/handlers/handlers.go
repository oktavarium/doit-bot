package handlers

import "github.com/gin-gonic/gin"

type Handlers struct {
	router *gin.Engine
}

func New(router *gin.Engine) *Handlers {
	h := &Handlers{
		router: router,
	}

	h.init()

	return h
}

func (h *Handlers) init() {
	h.router.POST("/", h.CreateData)
	h.router.GET("/", h.GetData)
}
