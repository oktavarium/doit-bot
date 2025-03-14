package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/domain/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *db) GetUserByTgId(ctx context.Context, tg_id int64) (*users.User, error) {
	var result dbo.User
	if err := db.users.FindOne(ctx, bson.M{"tg_id": tg_id}).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, users.ErrUserNotFound
		}
		return nil, fmt.Errorf("find user: %w", err)
	}

	return result.ToDomainUser()
}
