package storage

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *db) UpdateList(ctx context.Context, list *planner.List) error {
	dboList := dbo.FromDomainList(list)
	filter := bson.M{"id": dboList.Id}
	if _, err := db.lists.ReplaceOne(ctx, filter, dboList); err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return planner.ErrNotFound
		default:
			return errors.Join(planner.ErrInfrastructureError, err)
		}
	}

	return nil
}
