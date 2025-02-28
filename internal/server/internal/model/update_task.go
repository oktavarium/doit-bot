package model

import (
	"context"
)

func (m *Model) UpdateTask(
	ctx context.Context,
	owner int64,
	id string,
	assignee *int64,
	summary *string,
	done *bool,
) error {
	return m.storage.UpdateTask(ctx, owner, id, assignee, summary, done)
}
