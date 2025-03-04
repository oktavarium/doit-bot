package mongodb

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/internal/storage/internal/mongodb/dbo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *db) CreateUser(
	ctx context.Context,
	tg_id int64,
	chat_tg_id int64,
	firstName string,
	lastName string,
	username string,
) (string, error) {
	user := dbo.User{
		TgId:      tg_id,
		ChatTgId:  chat_tg_id,
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
	}

	result, err := db.users.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("insert one: %w", err)
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("get inserted task id")
	}

	return id.Hex(), nil
}
