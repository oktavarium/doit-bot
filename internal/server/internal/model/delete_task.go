package model

import (
	"context"
)

func (m *Model) DeleteTaskById(ctx context.Context, owner int64, id string) error {
	return m.storage.DeleteTaskById(ctx, owner, id)
}
