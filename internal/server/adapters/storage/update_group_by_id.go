package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *db) UpdateGroupById(
	ctx context.Context,
	actorId string,
	chatId string,
	name string,
) error {
	bsonUserId, err := primitive.ObjectIDFromHex(actorId)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	bsonGroupId, err := primitive.ObjectIDFromHex(chatId)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	filter := bson.M{"_id": bsonGroupId}
	update := bson.M{"$set": bson.M{"name": name}, "$addToSet": bson.M{"users": bsonUserId}}
	_, err = db.users.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("update one: %w", err)
	}

	return nil
}
