package logger

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/common"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		slog.Info(
			"new request",
			slog.String("path", c.FullPath()),
			slog.String("method", c.Request.Method),
			slog.String(common.HeaderAuthorization, c.GetHeader(common.HeaderAuthorization)),
		)

		c.Next()

		slog.Info(
			"request ended",
			slog.Int("status", c.Writer.Status()),
		)
	}
}
