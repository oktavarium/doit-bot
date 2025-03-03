package storage

import (
	"context"

	"github.com/oktavarium/doit-bot/internal/server/internal/dto"
)

type Storage interface {
	CreateTask(
		ctx context.Context,
		actorId string,
		assigneeId *string,
		listId *string,
		summary string,
		description *string,
	) (id string, err error)
	UpdateTaskById(ctx context.Context,
		actorId string,
		taskId string,
		assigneeId *string,
		summary *string,
		description *string,
		done *bool,
	) error
	SetTaskDoneById(ctx context.Context,
		actorId string,
		taskId string,
		done bool,
	) error
	GetTaskById(ctx context.Context, taskId string) (*dto.Task, error)
	DeleteTaskById(ctx context.Context, actorId string, taskId string) error
	GetTasksByOwner(ctx context.Context, actorId string) ([]*dto.Task, error)
	GetUserByTgId(ctx context.Context, tg_id int64) (*dto.User, error)
}
