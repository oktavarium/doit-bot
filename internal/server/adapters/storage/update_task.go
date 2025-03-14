package storage

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *db) UpdateTask(ctx context.Context, task *planner.Task) error {
	dboTask := dbo.FromDomainTask(task)
	filter := bson.M{"_id": dboTask.DbId}
	update := bson.M{"$set": dboTask}
	if _, err := db.tasks.UpdateOne(ctx, filter, update); err != nil {
		return fmt.Errorf("update one: %w", err)
	}

	return nil
}
