package common

import (
	"context"

	initdata "github.com/telegram-mini-apps/init-data-golang"
)

type contextKey string

const (
	initDataKey contextKey = "init-data"
	actorIdKey  contextKey = "actor-id"
)

func InitDataToContext(ctx context.Context, initData initdata.InitData) context.Context {
	return context.WithValue(ctx, initDataKey, initData)
}

func InitDataFromContext(ctx context.Context) (initdata.InitData, bool) {
	initData, ok := ctx.Value(initDataKey).(initdata.InitData)
	return initData, ok
}

func ActorIdToContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, actorIdKey, id)
}

func ActorIdFromContext(ctx context.Context) (string, bool) {
	initData, ok := ctx.Value(actorIdKey).(string)
	return initData, ok
}
