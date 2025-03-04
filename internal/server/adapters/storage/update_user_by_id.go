package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *db) UpdateUserById(
	ctx context.Context,
	actorId string,
	tg_id int64,
	chat_tg_id int64,
	firstName string,
	lastName string,
	username string,
) error {
	bsonActorId, err := primitive.ObjectIDFromHex(actorId)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	filter := bson.M{"_id": bsonActorId}
	updatePayload := bson.M{
		"tg_id":      tg_id,
		"chat_tg_id": chat_tg_id,
		"first_name": firstName,
		"last_name":  lastName,
		"username":   username,
	}

	update := bson.M{"$set": updatePayload}
	_, err = db.users.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("update one: %w", err)
	}

	return nil
}
