package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *storage) CreateTask(
	ctx context.Context,
	owner string,
	assignee *string,
	list *string,
	summary string,
	description *string,
) (string, error) {
	bsonOwnerId, err := primitive.ObjectIDFromHex(owner)
	if err != nil {
		return "", fmt.Errorf("invalid id: %w", err)
	}

	assigneeId, err := primitive.ObjectIDFromHex(assignee)
	if err != nil {
		return "", fmt.Errorf("invalid id: %w", err)
	}

	result, err := db.tasks.InsertOne(ctx, dbTask{
		Owner:    bsonOwnerId,
		Assignee: assignee,
		Summary:  &summary,
	})
	if err != nil {
		return "", fmt.Errorf("insert one: %w", err)
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("get inserted task id")
	}

	return id.Hex(), nil
}
