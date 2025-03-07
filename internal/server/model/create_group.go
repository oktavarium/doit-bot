package model

import (
	"context"
)

func (m *Model) CreateGroup(
	ctx context.Context,
	actorId string,
	name string,
) (string, error) {
	return m.storage.CreateGroup(ctx, actorId, name)
}
