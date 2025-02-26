package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/internal/api/internal/common"
	"github.com/oktavarium/doit-bot/internal/server/internal/dto"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (h *Handlers) CreateData(c *gin.Context) {
	initData := c.GetHeader(common.HeaderAuthorization)
	parsedData, err := initdata.Parse(initData)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var task dto.Task
	ownerID := strconv.Itoa(int(parsedData.User.ID))
	task.Owner = &ownerID
	if err := c.BindJSON(&task); err != nil {
		return
	}

	h.storage.CreateTask(c, &task)
}
