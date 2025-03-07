package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/adapters/httpapi/common"
)

func (h *Handlers) GetTasks(c *gin.Context) {
	actorId, ok := common.ActorIdFromContext(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, common.NewStatusResponse(http.StatusInternalServerError, ""))
		return
	}

	tasks, err := h.model.GetTasks(c, actorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewStatusResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, getTasksResponse{Tasks: tasks, Status: common.Status{Code: http.StatusOK}})
}
