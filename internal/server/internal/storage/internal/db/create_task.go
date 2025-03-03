package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *storage) CreateTask(
	ctx context.Context,
	actorId string,
	assigneeId *string,
	listId *string,
	summary string,
	description *string,
) (string, error) {
	bsonActorId, err := primitive.ObjectIDFromHex(actorId)
	if err != nil {
		return "", fmt.Errorf("invalid id: %w", err)
	}

	var bsonAssigneeId primitive.ObjectID
	if assigneeId != nil {
		bsonAssigneeId, err = primitive.ObjectIDFromHex(*assigneeId)
		if err != nil {
			return "", fmt.Errorf("invalid id: %w", err)
		}
	}

	task := dbTask{
		OwnerId:     bsonActorId,
		Summary:     summary,
		Description: *description,
	}
	if assigneeId != nil {
		task.AssigneeId = bsonAssigneeId
	}

	result, err := db.tasks.InsertOne(ctx, task)
	if err != nil {
		return "", fmt.Errorf("insert one: %w", err)
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("get inserted task id")
	}

	return id.Hex(), nil
}
