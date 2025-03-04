package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *db) UpdateTaskById(ctx context.Context,
	actorId string,
	taskId string,
	assigneeId *string,
	listId *string,
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

	var bsonListId primitive.ObjectID
	if listId != nil {
		bsonListId, err = primitive.ObjectIDFromHex(*listId)
		if err != nil {
			return fmt.Errorf("invalid id: %w", err)
		}
	}

	filter := bson.M{"_id": bsonTaskId, "owner": bsonActorId}
	updatePayload := bson.M{}
	if assigneeId != nil {
		updatePayload["assignee"] = bsonAssigneeId
	}
	if listId != nil {
		updatePayload["list_id"] = bsonListId
	}
	if summary != nil {
		updatePayload["summary"] = summary
	}
	if description != nil {
		updatePayload["description"] = description
	}
	if done != nil {
		updatePayload["done"] = done
	}

	update := bson.M{"$set": updatePayload}

	_, err = db.tasks.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("update one: %w", err)
	}

	return nil
}
