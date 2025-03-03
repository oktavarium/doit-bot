package db

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/internal/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *storage) GetTaskById(ctx context.Context, id string) (*dto.User, error) {
	var result dbUser
	bsonId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	if err := db.tasks.FindOne(ctx, bson.M{"_id": bsonId}).Decode(&result); err != nil {
		return nil, fmt.Errorf("find task: %w", err)
	}

	return dbUserToDTOUser(result), nil
}
