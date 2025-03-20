package adminquery

import (
	"context"
)

type IsAdmin struct {
	TgId int64
}

type isAdminHandler struct {
	adminsMap map[int64]struct{}
}

type IsAdminHandler QueryHandler[IsAdmin, bool]

func NewIsAdminHandler(admins []int64) IsAdminHandler {
	adminsMap := make(map[int64]struct{})
	for _, admin := range admins {
		adminsMap[admin] = struct{}{}
	}
	return isAdminHandler{
		adminsMap: adminsMap,
	}
}

func (h isAdminHandler) Handle(ctx context.Context, cmd IsAdmin) (bool, error) {
	_, ok := h.adminsMap[cmd.TgId]
	return ok, nil
}
