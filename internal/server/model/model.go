package model

import (
	"github.com/oktavarium/doit-bot/internal/server/ports"
)

type Model struct {
	tgclient ports.TGClient
	storage  ports.Storage
}

func New(tgclient ports.TGClient, storage ports.Storage) *Model {
	return &Model{
		tgclient: tgclient,
		storage:  storage,
	}
}
