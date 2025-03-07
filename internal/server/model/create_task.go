package model

import (
	"context"
)

func (m *Model) CreateTask(
	ctx context.Context,
	actorId string,
	assigneeId *string,
	listId *string,
	name string,
	description string,
) (string, error) {
	return m.storage.CreateTask(ctx, actorId, assigneeId, listId, name, description)
}
