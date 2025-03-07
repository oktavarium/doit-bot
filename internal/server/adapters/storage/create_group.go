package storage

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *db) CreateGroup(
	ctx context.Context,
	actorId string,
	name string,
) (string, error) {
	bsonActorId, err := primitive.ObjectIDFromHex(actorId)
	if err != nil {
		return "", fmt.Errorf("invalid id: %w", err)
	}

	group := dbo.Group{
		Users: []primitive.ObjectID{bsonActorId},
		Name:  name,
		Type:  dbo.WithoutChat,
	}

	result, err := db.groups.InsertOne(ctx, group)
	if err != nil {
		return "", fmt.Errorf("insert one: %w", err)
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("get inserted group id")
	}

	return id.Hex(), nil
}
