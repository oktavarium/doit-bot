package model

import (
	"context"
)

func (m *Model) UpdateTaskById(
	ctx context.Context,
	owner int64,
	id string,
	assignee *int64,
	summary *string,
	done *bool,
) error {
	return m.storage.UpdateTaskById(ctx, owner, id, assignee, summary, done)
}
