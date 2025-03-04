package model

import "context"

func (m *Model) SetTaskDoneById(
	ctx context.Context,
	actorId string,
	taskId string,
	done bool,
) error {
	return m.storage.SetTaskDoneById(ctx, actorId, taskId, done)
}
