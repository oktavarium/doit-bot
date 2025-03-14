package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (db *db) DeleteTask(ctx context.Context, actorId string, taskId string) error {
	filter := bson.M{"id": taskId, "owner_id": actorId}

	if _, err := db.tasks.DeleteOne(ctx, filter); err != nil {
		return fmt.Errorf("delete one: %w", err)
	}

	return nil
}
