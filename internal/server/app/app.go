package app

import (
	"github.com/oktavarium/doit-bot/internal/server/adapters"
	"github.com/oktavarium/doit-bot/internal/server/app/command"
	"github.com/oktavarium/doit-bot/internal/server/app/query"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
	"github.com/oktavarium/doit-bot/internal/server/domain/users"
)

type App struct {
	tgclient adapters.TGClient
	Commands commands
	Queries  queries
}

type commands struct {
	CreateTask    command.CreateTaskHandler
	CreateUser    command.CreateUserHandler
	DeleteTask    command.DeleteTaskHandler
	SetTaskStatus command.SetTaskStatusHandler
}

type queries struct {
	GetTasks      query.GetTasksHandler
	GetUserByTgId query.GetUserByTgIdHandler
	GetTask       query.GetTaskHandler
}

func New(
	tgclient adapters.TGClient,
	plannerDomainSerice *planner.DomainService,
	usersDomainService *users.DomainService,
) *App {
	return &App{
		tgclient: tgclient,
		Commands: commands{
			CreateTask:    command.NewCreateTaskHandler(plannerDomainSerice),
			CreateUser:    command.NewCreateUserHandler(usersDomainService),
			DeleteTask:    command.NewDeleteTaskHandler(plannerDomainSerice),
			SetTaskStatus: command.NewSetTaskStatusHandler(plannerDomainSerice),
		},
		Queries: queries{
			GetTasks:      query.NewGetTaskskHandler(plannerDomainSerice),
			GetUserByTgId: query.NewGetUserByTgIdHandler(usersDomainService),
			GetTask:       query.NewGetTaskHandler(plannerDomainSerice),
		},
	}
}
