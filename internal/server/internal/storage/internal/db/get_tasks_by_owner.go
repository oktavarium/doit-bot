package db

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/internal/dto"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *storage) GetTasksByOwner(ctx context.Context, owner int64) ([]*dto.Task, error) {
	cursor, err := db.tasks.Find(ctx, bson.M{"owner": owner})
	if err != nil {
		return nil, fmt.Errorf("find tasks: %w", err)
	}

	var tasks []dbTask
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, fmt.Errorf("read cursor: %w", err)
	}

	result := make([]*dto.Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, dbTaskToDtoTask(task))
	}

	return result, nil
}
