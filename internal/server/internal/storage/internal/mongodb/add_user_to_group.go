package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	filter := bson.M{"_id": bsonGroupId}
	update := bson.M{"$addToSet": bson.M{"users": bsonUserId}}

	_, err = db.groups.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("update one: %w", err)
	}

	return nil
}
