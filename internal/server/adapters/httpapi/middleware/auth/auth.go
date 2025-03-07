package auth

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/doiterr"
	"github.com/oktavarium/doit-bot/internal/server/adapters/httpapi/common"
	"github.com/oktavarium/doit-bot/internal/server/ports"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func Middleware(token string, m ports.Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		// We expect passing init data in the Authorization header in the following format:
		// <auth-type> <auth-data>
		// <auth-type> must be "tma", and <auth-data> is Telegram Mini Apps init data.
		authParts := strings.Split(c.GetHeader(common.HeaderAuthorization), " ")
		if len(authParts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		authType := authParts[0]
		authData := authParts[1]
		switch authType {
		case common.AuthTypeDebug:
			userTgId, err := strconv.Atoi(authData)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusEarlyHints, gin.H{
					"message": "Send me good debug auth header id you use dbg. For example: dbg 11223344",
				})
				return
			}

			userId, err := m.GetUserIdByTgId(c, int64(userTgId))
			if err != nil && !errors.Is(err, doiterr.ErrNotFound) {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			}

			if errors.Is(err, doiterr.ErrNotFound) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": err.Error(),
				})
				return
			}

			ctx := common.ActorIdToContext(c.Request.Context(), userId)
			c.Request = c.Request.WithContext(ctx)
		case common.AuthTypeTelegram:
			// Validate init data. We consider init data sign valid for 1 hour from their
			// creation moment.
			if err := initdata.Validate(authData, token, time.Hour); err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": err.Error(),
				})
				return
			}

			// Parse init data. We will surely need it in the future.
			initData, err := initdata.Parse(authData)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
				return
			}

			userId, err := m.GetUserIdByTgId(c, initData.User.ID)
			if err != nil && !errors.Is(err, doiterr.ErrNotFound) {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			}

			if errors.Is(err, doiterr.ErrNotFound) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": err.Error(),
				})
				return
			}

			ctx := common.InitDataToContext(c.Request.Context(), initData)
			ctx = common.ActorIdToContext(ctx, userId)
			c.Request = c.Request.WithContext(ctx)

		default:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "not supported auth type",
			})
		}
	}
}
