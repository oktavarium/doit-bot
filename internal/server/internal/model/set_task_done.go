package model

import "context"

func (m *Model) SetTaskDoneById(
	ctx context.Context,
	owner int64,
	id string,
	done bool,
) error {
	return m.storage.SetTaskDoneById(ctx, owner, id, done)
}
