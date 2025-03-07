package storage

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *db) GetTasks(ctx context.Context, actorId string) ([]*dto.Task, error) {
	bsonActorId, err := primitive.ObjectIDFromHex(actorId)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	filter := bson.M{"user_id": bsonActorId}
	cursor, err := db.utlinks.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("find utlinks: %w", err)
	}

	// TODO rewrite to lookup with condition
	var utLinks []dbo.UTLink
	if err = cursor.All(ctx, &utLinks); err != nil {
		return nil, fmt.Errorf("read cursor: %w", err)
	}

	taskIds := make([]primitive.ObjectID, 0, len(utLinks))
	for _, utLink := range utLinks {
		taskIds = append(taskIds, utLink.TaskId)
	}

	filter = bson.M{"_id": bson.M{"$in": taskIds}}
	cursor, err = db.tasks.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("find tasks: %w", err)
	}

	var tasks []dbo.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, fmt.Errorf("read cursor: %w", err)
	}
	result := make([]*dto.Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, task.ToDTOTask())
	}

	return result, nil
}
