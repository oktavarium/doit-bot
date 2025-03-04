package model

import (
	"context"
	"fmt"
)

func (m *Model) SendStartupButton(ctx context.Context, chatID int64, userID int64, username string) error {
	if err := m.tgclient.SendStartupButton(
		ctx,
		"Start WebAPP",
		"Hello guys! Someone added me to this psycho...",
		chatID,
	); err != nil {
		return fmt.Errorf("send startup message: %w", err)
	}
	return nil
}
