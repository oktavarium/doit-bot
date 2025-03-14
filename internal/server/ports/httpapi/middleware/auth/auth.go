package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/doiterr"
	"github.com/oktavarium/doit-bot/internal/server/app"
	"github.com/oktavarium/doit-bot/internal/server/app/query"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/common"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func Middleware(token string, app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// We expect passing init data in the Authorization header in the following format:
		// <auth-type> <auth-data>
		// <auth-type> must be "tma", and <auth-data> is Telegram Mini Apps init data.
		authParts := strings.Split(c.GetHeader(common.HeaderAuthorization), " ")
		if len(authParts) != 2 {
			c.AbortWithStatusJSON(http.StatusBadRequest, common.NewStatusResponse(http.StatusBadRequest, "wrong authorization scheme"))
			return
		}

		authType := authParts[0]
		authData := authParts[1]
		switch authType {
		case common.AuthTypeTelegram:
			// Validate init data. We consider init data sign valid for 1 hour from their
			// creation moment.
			if err := initdata.Validate(authData, token, time.Hour); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, common.NewStatusResponse(http.StatusBadRequest, err.Error()))
				return
			}

			initData, err := initdata.Parse(authData)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, common.NewStatusResponse(http.StatusBadRequest, err.Error()))
				return
			}
			user, err := app.Queries.GetUserByTgId.Handle(c, query.GetUserByTgId{TgId: initData.User.ID})
			if err != nil && !errors.Is(err, doiterr.ErrNotFound) {
				c.AbortWithStatusJSON(http.StatusInternalServerError, common.NewStatusResponse(http.StatusInternalServerError, err.Error()))
				return
			}

			if errors.Is(err, doiterr.ErrNotFound) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, common.NewStatusResponse(http.StatusUnauthorized, "user is not registered"))
				return
			}

			ctx := common.InitDataToContext(c.Request.Context(), initData)
			ctx = common.ActorIdToContext(ctx, user.Id())
			c.Request = c.Request.WithContext(ctx)

		default:
			c.AbortWithStatusJSON(http.StatusBadRequest, common.NewStatusResponse(http.StatusBadRequest, "not supported authentication scheme"))
		}
	}
}
