package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *storage) DeleteTaskById(ctx context.Context, actorId string, taskId string) error {
	bsonActorId, err := primitive.ObjectIDFromHex(actorId)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	bsonTaskId, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	filter := bson.D{
		{
			Key:   "_id",
			Value: bsonTaskId,
		},
		{
			Key:   "owner",
			Value: bsonActorId,
		},
	}

	_, err = db.tasks.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("delete one: %w", err)
	}

	return nil
}
