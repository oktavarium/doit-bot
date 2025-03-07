package model

import (
	"context"
)

func (m *Model) UpdateTaskById(
	ctx context.Context,
	actorId string,
	taskId string,
	assigneeId *string,
	listId *string,
	name *string,
	description *string,
	done *bool,
) error {
	return m.storage.UpdateTaskById(ctx, actorId, taskId, assigneeId, listId, name, description, done)
}
