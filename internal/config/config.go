package config

import (
	"fmt"
	"os"

	"github.com/oktavarium/doit-bot/internal/doiterr"
)

type Config struct {
	token         string
	listenAddress string
	uri           string
}

func GetConfig() (*Config, error) {
	cfg := &Config{
		token:         os.Getenv("BOT_TOKEN"),
		listenAddress: os.Getenv("LISTEN_ADDRESS"),
		uri:           os.Getenv("DB_URI"),
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

	if c.GetUri() == "" {
		return doiterr.ErrEmptyDbURI
	}

	return nil
}

func (c *Config) GetToken() string {
	return c.token
}

func (c *Config) GetEndpoint() string {
	return c.listenAddress
}

func (c *Config) GetUri() string {
	return c.uri
}
