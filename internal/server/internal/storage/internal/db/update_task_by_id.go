package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *storage) UpdateTaskById(ctx context.Context,
	actorId string,
	taskId string,
	assigneeId *string,
	summary *string,
	description *string,
	done *bool,
) error {
	bsonActorId, err := primitive.ObjectIDFromHex(actorId)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	bsonTaskId, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	var bsonAssigneeId primitive.ObjectID
	if assigneeId != nil {
		bsonAssigneeId, err = primitive.ObjectIDFromHex(*assigneeId)
		if err != nil {
			return fmt.Errorf("invalid id: %w", err)
		}
	}

	filter := bson.D{{"_id", bsonTaskId}, {"owner", bsonActorId}}
	updatePayload := bson.D{}
	if assigneeId != nil {
		updatePayload = append(updatePayload, bson.E{"assignee", bsonAssigneeId})
	}
	if summary != nil {
		updatePayload = append(updatePayload, bson.E{"summary", summary})
	}
	if description != nil {
		updatePayload = append(updatePayload, bson.E{"description", description})
	}
	if done != nil {
		updatePayload = append(updatePayload, bson.E{"done", done})
	}

	update := bson.D{{"$set", updatePayload}}

	_, err = db.tasks.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("update one: %w", err)
	}

	return nil
}
