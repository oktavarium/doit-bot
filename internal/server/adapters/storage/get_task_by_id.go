package storage

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/doiterr"
	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *db) GetTask(ctx context.Context, actorId string, taskId string) (*planner.Task, error) {
	var result dbo.Task
	filter := bson.M{"id": taskId, "owner_id": actorId}
	if err := db.tasks.FindOne(ctx, filter).Decode(&result); err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, planner.ErrTaskNotFound
		default:
			return nil, doiterr.WrapError(planner.ErrInfrastructureError, err)
		}
	}

	return result.ToDomainTask()
}
