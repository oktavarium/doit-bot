package middleware

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/app"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/middleware/auth"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/middleware/cors"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/middleware/logger"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/middleware/ratelimiter"
	"github.com/oktavarium/doit-bot/web"
)

const maxBodySize = 1000 // bytes

func Init(router *gin.Engine, token string, app *app.App) {
	router.Use(cors.Middleware()) // Disable after testing
	router.Use(ratelimiter.Middleware())
	router.Use(auth.Middleware(token, app))
	router.Use(logger.Middleware())
	router.Use(gin.Recovery())

	router.Use(static.Serve("/", static.EmbedFolder(web.StaticFiles, "client/build")))
}
