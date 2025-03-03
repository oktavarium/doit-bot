package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *storage) UpdateTaskById(ctx context.Context,
	owner int64,
	id string,
	assignee *int64,
	summary *string,
	done *bool,
) error {
	collection := db.client.Database("test").Collection("trainers")

	bsonId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	filter := bson.D{{"_id", bsonId}, {"owner", owner}}
	updatePayload := bson.D{}
	if assignee != nil {
		updatePayload = append(updatePayload, bson.E{"assignee", assignee})
	}
	if summary != nil {
		updatePayload = append(updatePayload, bson.E{"summary", summary})
	}
	if done != nil {
		updatePayload = append(updatePayload, bson.E{"done", done})
	}

	update := bson.D{{"$set", updatePayload}}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("update one: %w", err)
	}

	return nil
}
