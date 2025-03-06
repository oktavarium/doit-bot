package model

import (
	"context"

	"github.com/oktavarium/doit-bot/internal/server/dto"
)

func (m *Model) GetGroups(ctx context.Context, actorId string) ([]*dto.Group, error) {
	return m.storage.GetGroups(ctx, actorId)
}
