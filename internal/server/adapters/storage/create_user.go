package storage

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/domain/users"
)

func (db *db) CreateUser(ctx context.Context, user *users.User) error {
	dboUser := dbo.FromDomainUser(user)

	if _, err := db.users.InsertOne(ctx, dboUser); err != nil {
		return fmt.Errorf("insert one: %w", err)
	}

	return nil
}
