package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/internal/api/internal/common"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (h *Handlers) GetTasks(c *gin.Context) {
	initData := c.GetHeader(common.HeaderAuthorization)
	parsedData, err := initdata.Parse(initData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks, err := h.model.GetTasks(c, parsedData.User.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getTasksResponse{tasks})
}
