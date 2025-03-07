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
		Name:  name,
		Type:  dbo.WithoutChat,
	}

	result, err := db.groups.InsertOne(ctx, group)
	if err != nil {
		return "", fmt.Errorf("insert one: %w", err)
	}

	groupId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("get inserted group id")
	}

	ugLink := dbo.UGLink{
		UserId:  bsonActorId,
		GroupId: groupId,
	}

	if _, err := db.uglinks.InsertOne(ctx, ugLink); err != nil {
		return "", fmt.Errorf("insert one: %w", err)
	}

	return groupId.Hex(), nil
}
