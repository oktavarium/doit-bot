package planner

import (
	"context"
)

type PlannerRepository interface {
	CreateTask(ctx context.Context, task *Task) error
	GetTasks(ctx context.Context, actorId string) ([]*Task, error)
	GetListTasks(ctx context.Context, actorId string, listId string) ([]*Task, error)
	GetTask(ctx context.Context, actorId string, taskId string) (*Task, error)
	DeleteTask(ctx context.Context, actorId string, taskId string) error
	UpdateTask(ctx context.Context, task *Task) error

	CreateList(ctx context.Context, list *List) error
	GetLists(ctx context.Context, actorId string) ([]*List, error)
	GetList(ctx context.Context, actorId string, listId string) (*List, error)
	DeleteList(ctx context.Context, actorId string, listId string) error
	UpdateList(ctx context.Context, list *List) error
}
