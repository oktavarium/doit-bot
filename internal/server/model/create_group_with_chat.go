package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/doiterr"
)

func (m *Model) CreateGroupWithChat(
	ctx context.Context,
	actorId string,
	chat_tg_id int64,
	name string,
) (string, error) {
	_, err := m.storage.GetGroupByTgId(ctx, chat_tg_id)
	if err != nil && !errors.Is(err, doiterr.ErrNotFound) {
		return "", fmt.Errorf("get group by tg id: %w", err)
	}

	if err == nil {
		return "", fmt.Errorf("group with this chat alreay exist")
	}

	chatId, err := m.storage.CreateGroupWithChat(ctx, actorId, chat_tg_id, name)
	if err != nil {
		return "", fmt.Errorf("create user: %w", err)
	}

	return chatId, nil
}
