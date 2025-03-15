package storage

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/doiterr"
	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/domain/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *db) GetUserByTgId(ctx context.Context, tg_id int64) (*users.User, error) {
	var result dbo.User
	filter := bson.M{"tg_id": tg_id}
	if err := db.users.FindOne(ctx, filter).Decode(&result); err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, users.ErrUserNotFound
		default:
			return nil, doiterr.WrapError(users.ErrInfrastructureError, err)
		}
	}

	return result.ToDomainUser()
}
