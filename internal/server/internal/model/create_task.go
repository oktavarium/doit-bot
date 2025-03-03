package model

import (
	"context"
	"fmt"
)

func (m *Model) CreateTask(
	ctx context.Context,
	owner_tg_id int64,
	assignee *string,
	list *string,
	summary string,
	description *string,
) (string, error) {
	ownerUser, err := m.storage.GetUserByTgId(ctx, owner_tg_id)
	if err != nil {
		return "", fmt.Errorf("get user by tg id: %w", err)
	}
	return m.storage.CreateTask(ctx, ownerUser.Id, assignee, list, summary, description)
}
