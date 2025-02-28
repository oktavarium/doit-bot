package model

import (
	"github.com/oktavarium/doit-bot/internal/server/internal/storage"
	"github.com/oktavarium/doit-bot/internal/server/internal/tg/tgclient"
)

type Model struct {
	tgclient *tgclient.TGClient
	storage  storage.Storage
}

func New(tgclient *tgclient.TGClient, storage storage.Storage) *Model {
	return &Model{
		tgclient: tgclient,
		storage:  storage,
	}
}
