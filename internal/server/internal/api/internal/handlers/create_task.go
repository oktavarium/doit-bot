package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/internal/api/internal/common"
)

func (h *Handlers) CreateTask(c *gin.Context) {
	actorId, ok := common.ActorIdFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var request createTaskRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateCreateTaskRequest(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskID, err := h.model.CreateTask(
		c,
		actorId,
		request.AssigneeId,
		request.ListId,
		request.Summary,
		request.Description,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createTaskResponse{Id: taskID})
}
