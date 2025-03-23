package storage

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
)

func (db *db) CreateList(ctx context.Context, list *planner.List) error {
	dboList := dbo.FromDomainList(list)

	if _, err := db.lists.InsertOne(ctx, dboList); err != nil {
		return errors.Join(planner.ErrInfrastructureError, err)
	}

	return nil
}
