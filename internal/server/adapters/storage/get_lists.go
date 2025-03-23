package storage

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *db) GetLists(ctx context.Context, actorId string) ([]*planner.List, error) {
	filter := bson.M{"owner_id": actorId}
	cursor, err := db.lists.Find(ctx, filter)
	if err != nil {
		return nil, errors.Join(planner.ErrInfrastructureError, err)
	}

	var lists []dbo.List
	if err = cursor.All(ctx, &lists); err != nil {
		return nil, errors.Join(planner.ErrInfrastructureError, err)
	}

	result, err := dbo.ListsToDomainLists(lists)
	if err != nil {
		return nil, err
	}

	return result, nil
}
