package auth

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/internal/api/internal/common"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func Middleware(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		initData := c.GetHeader(common.HeaderAuthorization)
		if initData == "" {
			slog.Info("bad request")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		expIn := 24 * time.Hour
		if _, err := initdata.Parse(initData); err != nil {
			slog.Info("bad request")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err := initdata.Validate(initData, token, expIn); err != nil {
			slog.Info("unauthorized")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
