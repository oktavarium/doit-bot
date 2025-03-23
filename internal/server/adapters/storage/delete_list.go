package storage

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *db) DeleteList(ctx context.Context, actorId string, listId string) error {
	filter := bson.M{"id": listId, "owner_id": actorId}

	if _, err := db.lists.DeleteOne(ctx, filter); err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return planner.ErrNotFound
		default:
			return errors.Join(planner.ErrInfrastructureError, err)
		}
	}

	return nil
}
