package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/adapters/http_api/common"
)

func (h *Handlers) DeleteTaskById(c *gin.Context) {
	actorId, ok := common.ActorIdFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var request deleteTaskByIdRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.model.DeleteTaskById(
		c,
		actorId,
		request.Id,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
