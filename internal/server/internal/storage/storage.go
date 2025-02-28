package storage

import "github.com/oktavarium/doit-bot/internal/server/internal/storage/internal/db"

func New(uri string) (Storage, error) {
	return db.New(uri)
}
