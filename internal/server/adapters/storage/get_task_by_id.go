package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/doiterr"
	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *db) GetTask(ctx context.Context, taskId string) (*planner.Task, error) {
	var result dbo.Task
	if err := db.users.FindOne(ctx, bson.M{"id": taskId}).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, doiterr.ErrNotFound
		}
		return nil, fmt.Errorf("find task: %w", err)
	}

	return result.ToDomainTask()
}
