package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/internal/api/internal/common"
)

func (h *Handlers) DeleteTask(c *gin.Context) {
	initData, ok := common.CtxInitData(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
			"message": "Init data not found",
		})
		return
	}

	var request deleteTaskRequest
	owner := initData.User.ID
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.model.DeleteTask(
		c,
		owner,
		request.Id,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
