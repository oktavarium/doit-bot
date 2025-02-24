package config

import (
	"fmt"
	"os"

	"github.com/oktavarium/doit-bot/internal/doiterr"
)

type Config struct {
	token    string
	endpoint string
}

func GetConfig() (*Config, error) {
	cfg := &Config{
		token:    os.Getenv("BOT_TOKEN"),
		endpoint: os.Getenv("ENDPOINT"),
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
		return doiterr.ErrEmptyEndpoint
	}

	return nil
}

func (c *Config) GetToken() string {
	return c.token
}

func (c *Config) GetEndpoint() string {
	return c.endpoint
}
