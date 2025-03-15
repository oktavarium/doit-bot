package error

import (
	"errors"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/common"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors[0]

			slog.Info(
				"error handling request",
				slog.Int("status", c.Writer.Status()),
				slog.String("error", err.Error()),
			)

			var Error common.Error
			if errors.As(err, &Error) {
				c.JSON(Error.Status, common.NewStatusResponse(Error.Status, Error.Error()))
			} else {
				c.JSON(c.Writer.Status(), common.NewStatusResponse(c.Writer.Status(), err.Error()))
			}
		}
	}
}
