package ports

import (
	"context"

	"github.com/oktavarium/doit-bot/internal/server/dto"
)

type Model interface {
	CreateGroup(
		ctx context.Context,
		actorId string,
		chat_tg_id int64,
		name string,
	) (string, error)
	CreateTask(
		ctx context.Context,
		actorId string,
		assigneeId *string,
		listId *string,
		summary string,
		description string,
	) (string, error)
	CreateUser(
		ctx context.Context,
		actor_tg_id int64,
		chat_tg_id int64,
		firstName string,
		lastName string,
		username string,
	) error
	DeleteTaskById(ctx context.Context, actorId string, id string) error
	GetTaskById(ctx context.Context, id string) (*dto.Task, error)
	GetTasksByOwner(ctx context.Context, actorId string) ([]*dto.Task, error)
	GetUserIdByTgId(ctx context.Context, id int64) (string, error)
	SendStartupButton(ctx context.Context, chatID int64, userID int64, username string) error
	SetTaskDoneById(
		ctx context.Context,
		actorId string,
		taskId string,
		done bool,
	) error
	UpdateTaskById(
		ctx context.Context,
		actorId string,
		taskId string,
		assigneeId *string,
		listId *string,
		summary *string,
		description *string,
		done *bool,
	) error
}
