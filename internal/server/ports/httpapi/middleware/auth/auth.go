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
				common.AbortContextWithError(c, common.NewBadRequestError(err))
				return
			}

			initData, err := initdata.Parse(authData)
			if err != nil {
				common.AbortContextWithError(c, common.NewBadRequestError(err))
				return
			}
			user, err := app.Queries.GetUserByTgId.Handle(c, query.GetUserByTgId{TgId: initData.User.ID})
			if err != nil && !errors.Is(err, doiterr.ErrNotFound) {
				common.AbortContextWithError(c, common.NewUnauthorizedError(err))
				return
			}

			if errors.Is(err, doiterr.ErrNotFound) {
				common.AbortContextWithError(c, common.NewUnauthorizedError(errors.New("user is not registered")))
				return
			}

			ctx := common.InitDataToContext(c.Request.Context(), initData)
			ctx = common.ActorIdToContext(ctx, user.Id())
			c.Request = c.Request.WithContext(ctx)

		default:
			common.AbortContextWithError(c, common.NewBadRequestError(errors.New("not supported authentication scheme")))
			return
		}
	}
}
