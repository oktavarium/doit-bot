package storage

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *db) GetGroupUsers(ctx context.Context, groupId string) ([]*dto.User, error) {
	bsonGroupId, err := primitive.ObjectIDFromHex(groupId)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	filter := bson.M{"group_id": bsonGroupId}
	cursor, err := db.uglinks.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("find uglinks: %w", err)
	}

	// TODO rewrite to lookup with condition
	var uglinks []dbo.UGLink
	if err = cursor.All(ctx, &uglinks); err != nil {
		return nil, fmt.Errorf("read cursor: %w", err)
	}

	userIds := make([]primitive.ObjectID, 0, len(uglinks))
	for _, uglink := range uglinks {
		userIds = append(userIds, uglink.UserId)
	}

	filter = bson.M{"_id": bson.M{"$in": userIds}}
	cursor, err = db.groups.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("find users: %w", err)
	}

	var users []dbo.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, fmt.Errorf("read cursor: %w", err)
	}
	result := make([]*dto.User, 0, len(users))
	for _, user := range users {
		result = append(result, user.ToDTOUser())
	}

	return result, nil
}
