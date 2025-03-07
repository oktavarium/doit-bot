package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/adapters/httpapi/common"
)

func (h *Handlers) CreateTask(c *gin.Context) {
	actorId, ok := common.ActorIdFromContext(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, common.NewStatusResponse(http.StatusInternalServerError, ""))
		return
	}

	var request createTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, common.NewStatusResponse(http.StatusBadRequest, err.Error()))
		return
	}

	taskID, err := h.model.CreateTask(
		c,
		actorId,
		request.AssigneeId,
		request.ListId,
		request.Name,
		request.Description,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewStatusResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, createTaskResponse{Id: taskID, Status: common.Status{Code: http.StatusOK}})
}
