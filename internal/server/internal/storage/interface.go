package storage

import (
	"context"

	"github.com/oktavarium/doit-bot/internal/server/internal/dto"
)

type Storage interface {
	CreateTask(
		ctx context.Context,
		owner int64,
		summary string,
		assignee *int64,
	) (id string, err error)
	UpdateTask(ctx context.Context,
		owner int64,
		id string,
		assignee *int64,
		summary *string,
		done *bool,
	) error
	GetTask(ctx context.Context, id string) (*dto.Task, error)
	DeleteTask(ctx context.Context, owner int64, id string) error
	GetTasks(ctx context.Context, owner int64) ([]*dto.Task, error)
}
