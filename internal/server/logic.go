package server

import (
	"context"
	"fmt"
	"strconv"
)

func (s *server) initLogic() {
	s.tgAPI.SetBotAddedCallback(s.sendStartupButton)
}

func (s *server) sendStartupButton(ctx context.Context, chatID int64, userID int64, username string) error {
	if err := s.botAPI.SendStartupButton(
		ctx,
		"Start WebAPP",
		"Hello!",
		strconv.Itoa(int(chatID)),
		s.cfg.GetEndpoint(),
	); err != nil {
		return fmt.Errorf("send startup message: %w", err)
	}
	return nil
}
