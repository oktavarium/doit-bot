package storage

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *db) GetListsByGroupId(ctx context.Context, groupId string) ([]*dto.List, error) {
	bsonGroupId, err := primitive.ObjectIDFromHex(groupId)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	cursor, err := db.lists.Find(ctx, bson.M{"group_id": bsonGroupId})
	if err != nil {
		return nil, fmt.Errorf("find tasks: %w", err)
	}

	var lists []dbo.List
	if err := cursor.All(ctx, &lists); err != nil {
		return nil, fmt.Errorf("read cursor: %w", err)
	}

	result := make([]*dto.List, 0, len(lists))
	for _, list := range lists {
		result = append(result, list.ToDTOList())
	}

	return result, nil
}
