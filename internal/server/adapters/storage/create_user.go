package storage

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/domain/users"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *db) CreateUser(ctx context.Context, user *users.User) error {
	dboUser := dbo.FromDomainUser(user)

	if _, err := db.users.InsertOne(ctx, dboUser); err != nil {
		switch {
		case mongo.IsDuplicateKeyError(err):
			return users.ErrUserExists
		default:
			return errors.Join(users.ErrInfrastructureError, err)
		}
	}

	return nil
}
