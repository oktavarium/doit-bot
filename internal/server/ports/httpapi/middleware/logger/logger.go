package logger

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/common"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		slog.Debug(
			"new request",
			slog.String("path", c.Request.URL.RawPath),
			slog.String("mehtod", c.Request.Method),
			slog.String(common.HeaderAuthorization, c.GetHeader(common.HeaderAuthorization)),
		)

		c.Next()
	}
}
