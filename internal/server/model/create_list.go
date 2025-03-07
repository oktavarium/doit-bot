package model

import "context"

func (m *Model) CreateList(
	ctx context.Context,
	actorId string,
	groupId *string,
	name string,
	description string,
) (string, error) {
	return m.storage.CreateList(ctx, actorId, groupId, name, description)
}
