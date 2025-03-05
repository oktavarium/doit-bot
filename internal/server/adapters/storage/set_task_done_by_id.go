package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *db) SetTaskDoneById(ctx context.Context,
	actorId string,
	taskId string,
	done bool,
) error {
	bsonActorId, err := primitive.ObjectIDFromHex(actorId)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	bsonTaskId, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	filter := bson.M{"_id": bsonTaskId, "owner_id": bsonActorId}
	update := bson.M{"$set": bson.M{"done": done}}

	_, err = db.tasks.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("update one: %w", err)
	}

	return nil
}
