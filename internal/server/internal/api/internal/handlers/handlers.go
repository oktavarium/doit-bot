package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/internal/api/internal/middleware/auth"
	"github.com/oktavarium/doit-bot/internal/server/internal/model"
)

type Handlers struct {
	router *gin.Engine
	token  string
	model  *model.Model
}

func New(router *gin.Engine, token string, model *model.Model) *Handlers {
	h := &Handlers{
		router: router,
		token:  token,
		model:  model,
	}

	h.init()

	return h
}

func (h *Handlers) init() {
	apiGroup := h.router.Group("/api", auth.Middleware(h.token))
	apiGroup.POST("/create_task", h.CreateTask)
	apiGroup.POST("/delete_task", h.DeleteTask)
	apiGroup.POST("/update_task", h.UpdateTask)
	apiGroup.POST("/get_tasks", h.GetTasks)
	apiGroup.POST("/get_task", h.GetTask)
}
