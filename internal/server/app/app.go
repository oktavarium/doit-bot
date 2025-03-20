package app

import (
	"github.com/oktavarium/doit-bot/internal/server/app/command"
	"github.com/oktavarium/doit-bot/internal/server/app/query"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
	"github.com/oktavarium/doit-bot/internal/server/domain/users"
)

type App struct {
	Commands commands
	Queries  queries
}

type commands struct {
	CreateTask command.CreateTaskHandler
	CreateUser command.CreateUserHandler
	DeleteTask command.DeleteTaskHandler
	UpdateTask command.UpdateTaskHandler
}

type queries struct {
	GetTasks      query.GetTasksHandler
	GetUserByTgId query.GetUserByTgIdHandler
	GetTask       query.GetTaskHandler
}

func New(
	plannerDomainSerice planner.DomainService,
	usersDomainService users.DomainService,
) *App {
	return &App{
		Commands: commands{
			CreateTask: command.NewCreateTaskHandler(plannerDomainSerice),
			CreateUser: command.NewCreateUserHandler(usersDomainService),
			DeleteTask: command.NewDeleteTaskHandler(plannerDomainSerice),
			UpdateTask: command.NewUpdateTaskHandler(plannerDomainSerice),
		},
		Queries: queries{
			GetTasks:      query.NewGetTaskskHandler(plannerDomainSerice),
			GetUserByTgId: query.NewGetUserByTgIdHandler(usersDomainService),
			GetTask:       query.NewGetTaskHandler(plannerDomainSerice),
		},
	}
}
