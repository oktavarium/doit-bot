package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/doiterr"
)

func (m *Model) CreateUser(
	ctx context.Context,
	actor_tg_id int64,
	chat_tg_id int64,
	firstName string,
	lastName string,
	username string,
) error {
	user, err := m.storage.GetUserByTgId(ctx, actor_tg_id)
	if err != nil && !errors.Is(err, doiterr.ErrNotFound) {
		return fmt.Errorf("get user by tg id: %w", err)
	}

	if errors.Is(err, doiterr.ErrNotFound) {
		_, err := m.storage.CreateUser(ctx, actor_tg_id, chat_tg_id, firstName, lastName, username)
		if err != nil {
			return fmt.Errorf("create user: %w", err)
		}
		return nil
	}

	if err := m.storage.UpdateUserById(ctx, user.Id, actor_tg_id, chat_tg_id, firstName, lastName, username); err != nil {
		return fmt.Errorf("update user by id: %w", err)
	}

	return nil
}
