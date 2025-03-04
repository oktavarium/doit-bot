package mongodb

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
	bsonActorId, err := primitive.ObjectIDFromHex(actorId)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	bsonGroupId, err := primitive.ObjectIDFromHex(chatId)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	filter := bson.D{{"_id", bsonGroupId}}
	updatePayload := bson.D{
		bson.E{"owner_id", bsonActorId},
		bson.E{"name", name},
	}

	update := bson.D{{"$set", updatePayload}}
	_, err = db.users.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("update one: %w", err)
	}

	return nil
}
