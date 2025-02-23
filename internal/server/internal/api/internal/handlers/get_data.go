package handlers

import (
	"log/slog"
	"net/http"

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

	c.String(http.StatusOK, "this is a test string")

	slog.Info("Get data request: ", slog.Any("init data", parsedData))
}
