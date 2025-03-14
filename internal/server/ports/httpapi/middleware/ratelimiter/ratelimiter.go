package ratelimiter

import (
	"net/http"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/common"
)

const (
	limitPerTime = time.Second
)

func getHandlerLimitValue(c *gin.Context) string {
	return c.ClientIP() + c.Request.URL.RawPath + getAuthorization(c.Request)
}

func errorHandler(c *gin.Context, _ ratelimit.Info) {
	c.Status(http.StatusTooManyRequests)
}

func Middleware() func(c *gin.Context) {
	limiterStorage := ratelimit.InMemoryStore(
		&ratelimit.InMemoryOptions{
			Rate:  limitPerTime,
			Limit: uint(10),
		},
	)
	rateLimiter := ratelimit.RateLimiter(
		limiterStorage,
		&ratelimit.Options{
			KeyFunc:      getHandlerLimitValue,
			ErrorHandler: errorHandler,
		},
	)
	return rateLimiter
}

func getAuthorization(r *http.Request) string {
	h := r.Header.Get(common.HeaderAuthorization)
	if len(h) == 0 {
		return ""
	}

	return h
}
