package storage

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *db) CreateTask(
	ctx context.Context,
	actorId string,
	assigneeId *string,
	listId *string,
	name string,
	description string,
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

	task := dbo.Task{
		OwnerId:     bsonActorId,
		Name:        name,
		Description: description,
	}
	if assigneeId != nil {
		task.AssigneeId = bsonAssigneeId
	}

	result, err := db.tasks.InsertOne(ctx, task)
	if err != nil {
		return "", fmt.Errorf("insert one: %w", err)
	}

	taskId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("get inserted task id")
	}

	utLink := dbo.UTLink{
		UserId: bsonActorId,
		TaskId: taskId,
	}

	if _, err := db.utlinks.InsertOne(ctx, utLink); err != nil {
		return "", fmt.Errorf("insert one: %w", err)
	}

	return taskId.Hex(), nil
}
