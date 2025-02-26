package storage

import "github.com/oktavarium/doit-bot/internal/server/internal/storage/internal/file"

func NewFileStorage() Storage {
	return file.NewStorage()
}
