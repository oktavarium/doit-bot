package mongodb

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

	filter := bson.D{{"_id", bsonActorId}}
	updatePayload := bson.D{
		bson.E{"tg_id", tg_id},
		bson.E{"chat_tg_id", chat_tg_id},
		bson.E{"first_name", firstName},
		bson.E{"last_name", lastName},
		bson.E{"username", username},
	}

	update := bson.D{{"$set", updatePayload}}
	_, err = db.users.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("update one: %w", err)
	}

	return nil
}
