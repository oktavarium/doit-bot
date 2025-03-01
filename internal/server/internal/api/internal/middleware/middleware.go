package middleware

import (
	sizelimiter "github.com/gin-contrib/size"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/internal/api/internal/middleware/ratelimiter"
	"github.com/oktavarium/doit-bot/web"
)

const maxBodySize = 1000 // bytes

func Init(router *gin.Engine, token string) {
	router.Use(ratelimiter.Middleware())
	router.Use(sizelimiter.RequestSizeLimiter(maxBodySize))
	router.Use(gin.Recovery())

	router.Use(static.Serve("/", static.EmbedFolder(web.StaticFiles, "client/build")))
}
