package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/adapters/http_api/common"
)

func (h *Handlers) GetTasksByOwner(c *gin.Context) {
	actorId, ok := common.ActorIdFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	tasks, err := h.model.GetTasksByOwner(c, actorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getTasksByOwnerResponse{tasks})
}
