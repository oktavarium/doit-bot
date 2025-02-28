package model

import (
	"context"

	"github.com/oktavarium/doit-bot/internal/server/internal/dto"
)

func (m *Model) GetTasks(ctx context.Context, owner int64) ([]*dto.Task, error) {
	return m.storage.GetTasks(ctx, owner)
}
