package storage

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *db) GetTasks(ctx context.Context, actorId string) ([]*planner.Task, error) {
	filter := bson.M{"owner_id": actorId}
	cursor, err := db.tasks.Find(ctx, filter)
	if err != nil {
		return nil, errors.Join(planner.ErrInfrastructureError, err)
	}

	var tasks []dbo.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, errors.Join(planner.ErrInfrastructureError, err)
	}

	result, err := dbo.TasksToDomainTasks(tasks)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (db *db) GetListTasks(ctx context.Context, actorId string, listId string) ([]*planner.Task, error) {
	filter := bson.M{"owner_id": actorId, "list_id": listId}
	cursor, err := db.tasks.Find(ctx, filter)
	if err != nil {
		return nil, errors.Join(planner.ErrInfrastructureError, err)
	}

	var tasks []dbo.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, errors.Join(planner.ErrInfrastructureError, err)
	}

	result, err := dbo.TasksToDomainTasks(tasks)
	if err != nil {
		return nil, err
	}

	return result, nil
}
