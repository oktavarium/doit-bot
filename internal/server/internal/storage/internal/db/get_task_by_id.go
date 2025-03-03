package db

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/internal/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *storage) GetTaskById(ctx context.Context, taskId string) (*dto.Task, error) {
	var result dbTask
	bsonTaskId, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	if err := db.tasks.FindOne(ctx, bson.M{"_id": bsonTaskId}).Decode(&result); err != nil {
		return nil, fmt.Errorf("find task: %w", err)
	}

	return dbTaskToDTOTask(result), nil
}
