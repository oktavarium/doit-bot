package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/app"
)

type Handlers struct {
	router *gin.Engine
	token  string
	app    *app.App
}

func New(router *gin.Engine, token string, app *app.App) *Handlers {
	h := &Handlers{
		router: router,
		token:  token,
		app:    app,
	}

	RegisterHandlers(router, h)

	return h
}
