package middleware

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/internal/api/internal/middleware/ratelimiter"
	"github.com/oktavarium/doit-bot/web"
)

func Init(router *gin.Engine, token string) {
	router.Use(ratelimiter.Middleware())
	router.Use(gin.Recovery())

	router.Use(static.Serve("/", static.EmbedFolder(web.StaticFiles, "build")))
}
