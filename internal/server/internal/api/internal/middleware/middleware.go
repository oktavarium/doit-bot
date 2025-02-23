package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/internal/api/internal/middleware/ratelimiter"
)

func Init(router *gin.Engine, token string) {
	router.Use(ratelimiter.Middleware())
	router.Use(gin.Recovery())
}
