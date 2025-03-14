package storage

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *db) UpdateTask(ctx context.Context, actorId string, task *planner.Task) error {
	dboTask := dbo.FromDomainTask(task)
	filter := bson.M{"id": dboTask.Id, "owner_id": actorId}
	if _, err := db.tasks.ReplaceOne(ctx, filter, dboTask); err != nil {
		return fmt.Errorf("update one: %w", err)
	}

	return nil
}
