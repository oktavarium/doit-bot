package db

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/internal/dto"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *storage) GetUserByTgId(ctx context.Context, id int64) (*dto.User, error) {
	var result dbUser
	if err := db.users.FindOne(ctx, bson.M{"tg_id": id}).Decode(&result); err != nil {
		return nil, fmt.Errorf("find task: %w", err)
	}

	return dbUserToDTOUser(result), nil
}
