package tg_api

import (
	"github.com/go-telegram/bot"
	"github.com/oktavarium/doit-bot/internal/server/internal/tg_api/internal/handlers"
	"github.com/oktavarium/doit-bot/internal/server/internal/types"
)

type TgAPI struct {
	handlers *handlers.Handlers
}

func New() (*TgAPI, error) {
	return &TgAPI{
		handlers: handlers.New(),
	}, nil
}

func (api *TgAPI) GetDefaultHandler() bot.HandlerFunc {
	return api.handlers.DefaultHandler
}

func (api *TgAPI) SetBotAddedCallback(cb types.BotAddedCallBack) {
	api.handlers.SetBotAddedCallback(cb)
}
