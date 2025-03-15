package storage

import (
	"context"

	"github.com/oktavarium/doit-bot/internal/doiterr"
	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *db) GetTasks(ctx context.Context, actorId string) ([]*planner.Task, error) {
	filter := bson.M{"owner_id": actorId}
	cursor, err := db.tasks.Find(ctx, filter)
	if err != nil {
		return nil, doiterr.WrapError(planner.ErrInfrastructureError, err)
	}

	var tasks []dbo.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, doiterr.WrapError(planner.ErrInfrastructureError, err)
	}

	result, err := dbo.TasksToDomainTasks(tasks)
	if err != nil {
		return nil, err
	}

	return result, nil
}
