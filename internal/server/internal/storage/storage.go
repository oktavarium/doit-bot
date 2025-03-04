package storage

import "github.com/oktavarium/doit-bot/internal/server/internal/storage/internal/mongodb"

func New(uri string) (Storage, error) {
	return mongodb.New(uri)
}
