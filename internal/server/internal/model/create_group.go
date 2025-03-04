package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/doiterr"
)

func (m *Model) CreateGroup(
	ctx context.Context,
	actor_tg_id int64,
	chat_tg_id int64,
	name string,
) error {
	user, err := m.storage.GetUserByTgId(ctx, actor_tg_id)
	if err != nil && !errors.Is(err, doiterr.ErrNotFound) {
		return fmt.Errorf("get user by tg id: %w", err)
	}

	if errors.Is(err, doiterr.ErrNotFound) {
		return doiterr.ErrNotFound
	}

	group, err := m.storage.GetGroupByTgId(ctx, chat_tg_id)
	if err != nil && !errors.Is(err, doiterr.ErrNotFound) {
		return fmt.Errorf("get group by tg id: %w", err)
	}

	if errors.Is(err, doiterr.ErrNotFound) {
		_, err := m.storage.CreateGroup(ctx, user.Id, chat_tg_id, name)
		if err != nil {
			return fmt.Errorf("create user: %w", err)
		}
		return nil
	}

	if err := m.storage.UpdateGroupById(ctx, user.Id, group.Id, name); err != nil {
		return fmt.Errorf("update user by id: %w", err)
	}

	return nil
}
