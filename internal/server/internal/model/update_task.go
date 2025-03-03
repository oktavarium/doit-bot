package model

import (
	"context"
)

func (m *Model) UpdateTaskById(
	ctx context.Context,
	actorId string,
	taskId string,
	assigneeId *string,
	summary *string,
	description *string,
	done *bool,
) error {
	return m.storage.UpdateTaskById(ctx, actorId, taskId, assigneeId, summary, description, done)
}
