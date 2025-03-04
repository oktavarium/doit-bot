package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/doiterr"
)

func (m *Model) GetUserIdByTgId(ctx context.Context, id int64) (string, error) {
	user, err := m.storage.GetUserByTgId(ctx, id)
	if err != nil && !errors.Is(err, doiterr.ErrNotFound) {
		return "", fmt.Errorf("get user by tg id: %w", err)
	}

	if errors.Is(err, doiterr.ErrNotFound) {
		return "", doiterr.ErrNotFound
	}

	return user.Id, nil
}
