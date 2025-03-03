package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *storage) DeleteTaskById(ctx context.Context, owner int64, id string) error {
	bsonId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	filter := bson.D{{"_id", bsonId}, {"owner", owner}}

	_, err = db.tasks.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("delete one: %w", err)
	}

	return nil
}
