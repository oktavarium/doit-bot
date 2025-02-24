package config

import (
	"fmt"
	"os"

	"github.com/oktavarium/doit-bot/internal/doiterr"
)

const defaultListenAddress = "0.0.0.0:8080"

type Config struct {
	token         string
	listenAddress string
}

func GetConfig() (*Config, error) {
	cfg := &Config{
		token:         os.Getenv("BOT_TOKEN"),
		listenAddress: os.Getenv("LISTEN_ADDRESS"),
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("validate config: %w", err)
	}
	return cfg, nil
}

func (c *Config) validate() error {
	if c.GetToken() == "" {
		return doiterr.ErrEmptyToken
	}

	if c.GetEndpoint() == "" {
		c.listenAddress = defaultListenAddress
	}

	return nil
}

func (c *Config) GetToken() string {
	return c.token
}

func (c *Config) GetEndpoint() string {
	return c.listenAddress
}
