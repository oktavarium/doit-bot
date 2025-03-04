package middleware

import (
	sizelimiter "github.com/gin-contrib/size"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/adapters/http_api/middleware/cors"
	"github.com/oktavarium/doit-bot/internal/server/adapters/http_api/middleware/logger"
	"github.com/oktavarium/doit-bot/internal/server/adapters/http_api/middleware/ratelimiter"
	"github.com/oktavarium/doit-bot/web"
)

const maxBodySize = 1000 // bytes

func Init(router *gin.Engine, token string) {
	router.Use(cors.Middleware()) // Disable after testing
	router.Use(ratelimiter.Middleware())
	router.Use(logger.Middleware())
	router.Use(sizelimiter.RequestSizeLimiter(maxBodySize))
	router.Use(gin.Recovery())

	router.Use(static.Serve("/", static.EmbedFolder(web.StaticFiles, "client/build")))
}
