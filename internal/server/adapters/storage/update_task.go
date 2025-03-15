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

func (db *db) UpdateTask(ctx context.Context, actorId string, task *planner.Task) error {
	dboTask := dbo.FromDomainTask(task)
	filter := bson.M{"id": dboTask.Id, "owner_id": actorId}
	if _, err := db.tasks.ReplaceOne(ctx, filter, dboTask); err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return planner.ErrTaskNotFound
		default:
			return doiterr.WrapError(planner.ErrInfrastructureError, err)
		}
	}

	return nil
}
