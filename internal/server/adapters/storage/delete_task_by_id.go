package storage

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/doiterr"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *db) DeleteTask(ctx context.Context, actorId string, taskId string) error {
	filter := bson.M{"id": taskId, "owner_id": actorId}

	if _, err := db.tasks.DeleteOne(ctx, filter); err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return planner.ErrTaskNotFound
		default:
			return doiterr.WrapError(planner.ErrInfrastructureError, err)
		}
	}

	return nil
}
