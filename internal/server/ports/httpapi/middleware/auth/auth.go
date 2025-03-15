package auth

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/app"
	"github.com/oktavarium/doit-bot/internal/server/app/apperr"
	"github.com/oktavarium/doit-bot/internal/server/app/query"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/common"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

var (
	ErrWrongAuthScheme        = errors.New("wrong authorization scheme")
	ErrNotSupprotedAuthScheme = errors.New("not supported authentication scheme")
)

func Middleware(token string, app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		authParts := strings.Split(c.GetHeader(common.HeaderAuthorization), " ")
		if len(authParts) != 2 {
			common.AbortContextWithError(c, common.NewBadRequestError(ErrWrongAuthScheme))
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
			if err != nil {
				switch {
				case errors.Is(err, apperr.ErrValidationError):
					common.AbortContextWithError(c, common.NewBadRequestError(err))
					return
				case errors.Is(err, apperr.ErrNotFoundError):
					common.AbortContextWithError(c, common.NewUnauthorizedError(err))
					return
				default:
					common.AbortContextWithError(c, common.NewInternalServerError(err))
					return
				}
			}

			ctx := common.InitDataToContext(c.Request.Context(), initData)
			ctx = common.ActorIdToContext(ctx, user.Id())
			c.Request = c.Request.WithContext(ctx)

		default:
			common.AbortContextWithError(c, common.NewBadRequestError(ErrNotSupprotedAuthScheme))
			return
		}
	}
}
