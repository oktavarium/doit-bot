package planner

import (
	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/app"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/common"
)

type Planner struct {
	router *gin.Engine
	app    *app.App
}

func errorHandler(c *gin.Context, err error, code int) {
	common.ErrorToContext(c, common.NewBadRequestError(err))
}

func New(router *gin.Engine, app *app.App, middlwares ...MiddlewareFunc) *Planner {
	h := &Planner{
		router: router,
		app:    app,
	}

	options := GinServerOptions{
		Middlewares:  middlwares,
		ErrorHandler: errorHandler,
	}

	RegisterHandlersWithOptions(router, h, options)

	return h
}
