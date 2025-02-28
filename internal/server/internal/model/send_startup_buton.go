package model

import (
	"context"
	"fmt"
	"strconv"
)

func (m *Model) SendStartupButton(ctx context.Context, chatID int64, userID int64, username string) error {
	if err := m.tgclient.SendStartupButton(
		ctx,
		"Start WebAPP",
		"Hello!",
		strconv.Itoa(int(chatID)),
	); err != nil {
		return fmt.Errorf("send startup message: %w", err)
	}
	return nil
}
