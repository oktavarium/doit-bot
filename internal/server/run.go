package server

import (
	"fmt"

	"github.com/oktavarium/doit-bot/internal/config"
)

// Run - main bot function
func Run() error {
	config, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("get config: %w", err)
	}

	server, err := newServer(config)
	if err != nil {
		return fmt.Errorf("new server: %w", err)
	}

	server.initLogic()

	server.serve()
	return nil
}
