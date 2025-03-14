package planner

import (
	"context"
)

type PlannerRepository interface {
	CreateTask(ctx context.Context, task *Task) error
	GetTasks(ctx context.Context, actorId string) ([]*Task, error)
	GetTask(ctx context.Context, taskId string) (*Task, error)
	DeleteTask(ctx context.Context, taskId string) error
	UpdateTask(ctx context.Context, task *Task) error
}
