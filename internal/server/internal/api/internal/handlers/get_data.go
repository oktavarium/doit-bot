package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/internal/api/internal/common"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (h *Handlers) GetData(c *gin.Context) {
	initData := c.GetHeader(common.HeaderAuthorization)
	parsedData, err := initdata.Parse(initData)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	tasks, _ := h.storage.GetTasks(c, strconv.Itoa(int(parsedData.User.ID)))
	c.JSON(http.StatusOK, tasks)

	slog.Info("Get data request: ", slog.Any("init data", parsedData))
}
