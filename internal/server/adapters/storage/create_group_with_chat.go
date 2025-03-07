package storage

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *db) CreateGroupWithChat(
	ctx context.Context,
	actorId string,
	chat_tg_id int64,
	name string,
) (string, error) {
	bsonActorId, err := primitive.ObjectIDFromHex(actorId)
	if err != nil {
		return "", fmt.Errorf("invalid id: %w", err)
	}

	group := dbo.Group{
		TgId:  chat_tg_id,
		Users: []primitive.ObjectID{bsonActorId},
		Type:  dbo.WithChat,
		Name:  name,
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
