package storage

import (
	"context"

	"github.com/oktavarium/doit-bot/internal/server/internal/dto"
)

type Storage interface {
	CreateTask(ctx context.Context, task *dto.Task) (id string, err error)
	UpdateTask(ctx context.Context, task *dto.Task) error
	GetTask(ctx context.Context, owner, id string) (*dto.Task, error)
	DeleteTask(ctx context.Context, owner, id string) error
	GetTasks(ctx context.Context, owner string) ([]*dto.Task, error)
}
