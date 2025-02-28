package model

import (
	"context"
)

func (m *Model) CreateTask(
	ctx context.Context,
	owner int64,
	summary string,
	assignee *int64,
) (string, error) {
	return m.storage.CreateTask(ctx, owner, summary, assignee)
}
