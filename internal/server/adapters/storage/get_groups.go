package storage

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *db) GetGroups(ctx context.Context, actorId string) ([]*dto.Group, error) {
	bsonActorId, err := primitive.ObjectIDFromHex(actorId)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	filter := bson.M{"user_id": bsonActorId}
	cursor, err := db.uglinks.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("find uglinks: %w", err)
	}

	// TODO rewrite to lookup with condition
	var ugLinks []dbo.UGLink
	if err = cursor.All(ctx, &ugLinks); err != nil {
		return nil, fmt.Errorf("read cursor: %w", err)
	}

	groupIds := make([]primitive.ObjectID, 0, len(ugLinks))
	for _, ugLink := range ugLinks {
		groupIds = append(groupIds, ugLink.GroupId)
	}

	filter = bson.M{"_id": bson.M{"$in": groupIds}}
	cursor, err = db.groups.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("find groups: %w", err)
	}

	var groups []dbo.Group
	if err := cursor.All(ctx, &groups); err != nil {
		return nil, fmt.Errorf("read cursor: %w", err)
	}
	result := make([]*dto.Group, 0, len(groups))
	for _, group := range groups {
		result = append(result, group.ToDTOGroup())
	}

	return result, nil
}
