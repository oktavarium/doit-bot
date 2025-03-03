package storage

import (
	"context"

	"github.com/oktavarium/doit-bot/internal/server/internal/dto"
)

type Storage interface {
	CreateTask(
		ctx context.Context,
		owner string,
		assignee *string,
		list *string,
		summary string,
		description *string,
	) (id string, err error)
	UpdateTaskById(ctx context.Context,
		owner int64,
		id string,
		assignee *int64,
		summary *string,
		done *bool,
	) error
	SetTaskDoneById(ctx context.Context,
		owner int64,
		id string,
		done bool,
	) error
	GetTaskById(ctx context.Context, id string) (*dto.Task, error)
	DeleteTaskById(ctx context.Context, owner int64, id string) error
	GetTasksByOwner(ctx context.Context, owner int64) ([]*dto.Task, error)
	GetUserByTgId(ctx context.Context, id int64) (*dto.User, error)
}
