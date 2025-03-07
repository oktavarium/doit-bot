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

func (db *db) GetGroupById(ctx context.Context, actorId string, groupId string) (*dto.Group, error) {
	bsonActorId, err := primitive.ObjectIDFromHex(actorId)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	bsonGroupId, err := primitive.ObjectIDFromHex(groupId)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	var result dbo.Group
	if err := db.groups.FindOne(ctx, bson.M{"_id": bsonGroupId, "users": bson.M{"$in": bsonActorId}}).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, doiterr.ErrNotFound
		}
		return nil, fmt.Errorf("find task: %w", err)
	}

	return result.ToDTOGroup(), nil
}
