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
		description string,
	) (string, error)
	CreateUser(
		ctx context.Context,
		tg_id int64,
		chat_tg_id int64,
		firstName string,
		lastName string,
		username string,
	) (string, error)
	CreateGroup(
		ctx context.Context,
		actorId string,
		chat_tg_id int64,
		name string,
	) (string, error)
	UpdateUserById(
		ctx context.Context,
		actorId string,
		tg_id int64,
		chat_tg_id int64,
		firstName string,
		lastName string,
		username string,
	) error
	UpdateTaskById(ctx context.Context,
		actorId string,
		taskId string,
		assigneeId *string,
		listId *string,
		summary *string,
		description *string,
		done *bool,
	) error
	UpdateGroupById(
		ctx context.Context,
		actorId string,
		chatId string,
		name string,
	) error
	SetTaskDoneById(ctx context.Context,
		actorId string,
		taskId string,
		done bool,
	) error
	AddUserToGroup(
		ctx context.Context,
		userId string,
		groupId string,
	) error
	GetTaskById(ctx context.Context, taskId string) (*dto.Task, error)
	GetTasksByOwner(ctx context.Context, actorId string) ([]*dto.Task, error)
	GetUserByTgId(ctx context.Context, tg_id int64) (*dto.User, error)
	GetGroupByTgId(ctx context.Context, tg_id int64) (*dto.Group, error)
	DeleteTaskById(ctx context.Context, actorId string, taskId string) error
	RemoveUserFromGroup(
		ctx context.Context,
		userId string,
		groupId string,
	) error
}
