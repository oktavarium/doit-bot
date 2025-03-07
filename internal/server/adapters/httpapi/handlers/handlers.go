package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/adapters/httpapi/middleware/auth"
	"github.com/oktavarium/doit-bot/internal/server/ports"
)

type Handlers struct {
	router *gin.Engine
	token  string
	model  ports.Model
}

func New(router *gin.Engine, token string, model ports.Model) *Handlers {
	h := &Handlers{
		router: router,
		token:  token,
		model:  model,
	}

	h.init()

	return h
}

func (h *Handlers) init() {
	apiGroup := h.router.Group("/api")
	apiGroup.POST("/register", h.Register)
	apiGroup.POST("/create_task", h.CreateTask, auth.Middleware(h.token, h.model))
	apiGroup.POST("/create_group", h.CreateGroup, auth.Middleware(h.token, h.model))
	apiGroup.POST("/create_list", h.CreateList, auth.Middleware(h.token, h.model))

	apiGroup.POST("/get_tasks", h.GetTasks, auth.Middleware(h.token, h.model))
	apiGroup.POST("/get_tasks_by_owner", h.GetTasksByOwner, auth.Middleware(h.token, h.model))
	apiGroup.POST("/get_task_by_id", h.GetTaskById, auth.Middleware(h.token, h.model))
	apiGroup.POST("/get_groups", h.GetGroups, auth.Middleware(h.token, h.model))
	apiGroup.POST("/get_lists_by_group_id", h.GetListsByGroupId, auth.Middleware(h.token, h.model))

	apiGroup.POST("/delete_task_by_id", h.DeleteTaskById, auth.Middleware(h.token, h.model))
	apiGroup.POST("/update_task_by_id", h.UpdateTaskById, auth.Middleware(h.token, h.model))

	apiGroup.POST("/set_task_done_by_id", h.SetTaskDoneById, auth.Middleware(h.token, h.model))
}
