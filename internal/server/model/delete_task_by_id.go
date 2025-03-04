package model

import (
	"context"
)

func (m *Model) DeleteTaskById(ctx context.Context, actorId string, id string) error {
	return m.storage.DeleteTaskById(ctx, actorId, id)
}
