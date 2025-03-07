package storage

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *db) CreateList(
	ctx context.Context,
	actorId string,
	groupId *string,
	name string,
	description string,
) (string, error) {
	bsonActorId, err := primitive.ObjectIDFromHex(actorId)
	if err != nil {
		return "", fmt.Errorf("invalid id: %w", err)
	}

	var bsonGroupId primitive.ObjectID
	if groupId != nil {
		bsonGroupId, err = primitive.ObjectIDFromHex(*groupId)
		if err != nil {
			return "", fmt.Errorf("invalid id: %w", err)
		}
	}

	list := dbo.List{
		OwnerId:     bsonActorId,
		Name:        name,
		Description: description,
	}
	if groupId != nil {
		list.GroupId = bsonGroupId
	}

	result, err := db.lists.InsertOne(ctx, list)
	if err != nil {
		return "", fmt.Errorf("insert one: %w", err)
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("get inserted list id")
	}

	return id.Hex(), nil
}
