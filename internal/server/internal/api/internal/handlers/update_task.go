package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/internal/api/internal/common"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (h *Handlers) UpdateTask(c *gin.Context) {
	initData := c.GetHeader(common.HeaderAuthorization)
	parsedData, err := initdata.Parse(initData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var request updateTaskRequest
	owner := parsedData.User.ID
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.model.UpdateTask(
		c,
		owner,
		request.Id,
		request.Assignee,
		request.Summary,
		request.Done,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
