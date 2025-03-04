package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/doiterr"
	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *db) GetTaskById(ctx context.Context, taskId string) (*dto.Task, error) {
	var result dbo.Task
	bsonTaskId, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	if err := db.tasks.FindOne(ctx, bson.M{"_id": bsonTaskId}).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, doiterr.ErrNotFound
		}
		return nil, fmt.Errorf("find task: %w", err)
	}

	return result.ToDTOTask(), nil
}
