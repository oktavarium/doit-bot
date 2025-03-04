package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/doiterr"
	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *db) GetUserByTgId(ctx context.Context, tg_id int64) (*dto.User, error) {
	var result dbo.User
	if err := db.users.FindOne(ctx, bson.M{"tg_id": tg_id}).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, doiterr.ErrNotFound
		}
		return nil, fmt.Errorf("find task: %w", err)
	}

	return result.ToDTOUser(), nil
}
