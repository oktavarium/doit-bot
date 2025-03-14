package noroute

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/common"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, common.NewStatusResponse(http.StatusNotFound, "no such method"))
	}
}
