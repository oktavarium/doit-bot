package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *storage) SetTaskDoneById(ctx context.Context,
	owner int64,
	id string,
	done bool,
) error {
	bsonId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	filter := bson.D{{"_id", bsonId}, {"owner", owner}}
	updatePayload := bson.D{}
	updatePayload = append(updatePayload, bson.E{"done", done})

	update := bson.D{{"$set", updatePayload}}

	_, err = db.tasks.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("update one: %w", err)
	}

	return nil
}
