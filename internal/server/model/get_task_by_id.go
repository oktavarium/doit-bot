package model

import (
	"context"

	"github.com/oktavarium/doit-bot/internal/server/dto"
)

func (m *Model) GetTaskById(ctx context.Context, id string) (*dto.Task, error) {
	return m.storage.GetTaskById(ctx, id)
}
