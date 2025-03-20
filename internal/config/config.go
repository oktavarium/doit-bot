package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/oktavarium/doit-bot/internal/doiterr"
)

type Config struct {
	token         string
	listenAddress string
	uri           string
	admins        []int64
}

func GetConfig() (*Config, error) {
	adminsRaw := os.Getenv("ADMINS")
	adminsSlice := strings.Split(adminsRaw, ",")
	admins := make([]int64, 0, len(adminsSlice))
	for _, admin := range adminsSlice {
		id, err := strconv.ParseInt(admin, 10, 64)
		if err != nil {
			continue
		}
		admins = append(admins, id)
	}

	cfg := &Config{
		token:         os.Getenv("BOT_TOKEN"),
		listenAddress: os.Getenv("LISTEN_ADDRESS"),
		uri:           os.Getenv("DB_URI"),
		admins:        admins,
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

func (c *Config) GetAdmins() []int64 {
	return c.admins
}
