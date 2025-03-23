package storage

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *db) GetList(ctx context.Context, actorId string, listId string) (*planner.List, error) {
	var result dbo.List
	filter := bson.M{"id": listId, "owner_id": actorId}
	if err := db.lists.FindOne(ctx, filter).Decode(&result); err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, planner.ErrNotFound
		default:
			return nil, errors.Join(planner.ErrInfrastructureError, err)
		}
	}

	return result.ToDomainList()
}
