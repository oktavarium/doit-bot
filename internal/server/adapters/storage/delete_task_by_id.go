package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *db) DeleteTask(ctx context.Context, taskId string) error {
	bsonTaskId, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	filter := bson.D{
		{
			Key:   "id",
			Value: bsonTaskId,
		},
	}

	_, err = db.tasks.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("delete one: %w", err)
	}

	return nil
}
