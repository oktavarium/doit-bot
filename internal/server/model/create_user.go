package model

import (
	"context"
	"errors"
	"fmt"
	"strconv"

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
	// TODO we should check if user exist by storage, not by our handsTODO:
	_, err := m.storage.GetUserByTgId(ctx, actor_tg_id)
	if err != nil && !errors.Is(err, doiterr.ErrNotFound) {
		return fmt.Errorf("get user by tg id: %w", err)
	}

	if err == nil {
		return fmt.Errorf("user already exist")
	}

	if _, err := m.storage.CreateUser(ctx, actor_tg_id, chat_tg_id, firstName, lastName, username); err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	// TODO For Test - should be in another place
	if err := m.tgclient.SendMessage(ctx, "Hello, user... this is your tg_id... "+strconv.FormatInt(actor_tg_id, 10), chat_tg_id); err != nil {
		return fmt.Errorf("send message to user: %w", err)
	}

	return nil
}
