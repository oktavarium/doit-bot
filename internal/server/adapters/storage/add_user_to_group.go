package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *db) AddUserToGroup(
	ctx context.Context,
	userId string,
	groupId string,
) error {
	bsonUserId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	bsonGroupId, err := primitive.ObjectIDFromHex(groupId)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	var uglink dbo.UGLink
	if err := db.uglinks.FindOne(ctx, bson.M{"user_id": bsonUserId, "group_id": bsonGroupId}).Decode(&uglink); err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("find ug link: %w", err)
		}
	}

	uglink = dbo.UGLink{
		UserId:  bsonUserId,
		GroupId: bsonGroupId,
	}

	if _, err := db.uglinks.InsertOne(ctx, uglink); err != nil {
		return fmt.Errorf("insert one: %w", err)
	}

	return nil
}
