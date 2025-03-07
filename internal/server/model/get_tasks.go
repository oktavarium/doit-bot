package model

import (
	"context"

	"github.com/oktavarium/doit-bot/internal/server/dto"
)

func (m *Model) GetTasks(ctx context.Context, actorId string) ([]*dto.Task, error) {
	return m.storage.GetTasks(ctx, actorId)
}
