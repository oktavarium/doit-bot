package model

import (
	"context"
	"fmt"
)

func (m *Model) UpdateGroupById(
	ctx context.Context,
	actorId string,
	groupId string,
	name string,
) error {
	if err := m.storage.UpdateGroupById(ctx, actorId, groupId, name); err != nil {
		return fmt.Errorf("update group by id: %w", err)
	}

	return nil
}
