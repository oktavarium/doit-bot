package model

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/dto"
)

func (m *Model) GetListsByGroupId(ctx context.Context, actorId string, groupId string) ([]*dto.List, error) {
	_, err := m.storage.GetGroupById(ctx, actorId, groupId)
	if err != nil {
		return nil, fmt.Errorf("get group by id: %w", err)
	}
	return m.storage.GetListsByGroupId(ctx, groupId)
}
