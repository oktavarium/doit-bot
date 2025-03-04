package mongodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/doiterr"
	"github.com/oktavarium/doit-bot/internal/server/internal/dto"
	"github.com/oktavarium/doit-bot/internal/server/internal/storage/internal/mongodb/dbo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *db) GetGroupByTgId(ctx context.Context, tg_id int64) (*dto.Group, error) {
	var result dbo.Group
	if err := db.groups.FindOne(ctx, bson.M{"tg_id": tg_id}).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, doiterr.ErrNotFound
		}
		return nil, fmt.Errorf("find task: %w", err)
	}

	return result.ToDTOGroup(), nil
}
