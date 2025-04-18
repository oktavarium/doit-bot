package storage

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/dbo"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
)

func (db *db) CreateTask(ctx context.Context, task *planner.Task) error {
	dboTask := dbo.FromDomainTask(task)

	if _, err := db.tasks.InsertOne(ctx, dboTask); err != nil {
		return errors.Join(planner.ErrInfrastructureError, err)
	}

	return nil
}
