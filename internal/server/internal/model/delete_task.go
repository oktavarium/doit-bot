package model

import (
	"context"
)

func (m *Model) DeleteTask(ctx context.Context, owner int64, id string) error {
	return m.storage.DeleteTask(ctx, owner, id)
}
