package app

import (
	"github.com/oktavarium/doit-bot/internal/server/app/admincommand"
	"github.com/oktavarium/doit-bot/internal/server/app/adminquery"
	"github.com/oktavarium/doit-bot/internal/server/app/command"
	"github.com/oktavarium/doit-bot/internal/server/app/query"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
	"github.com/oktavarium/doit-bot/internal/server/domain/users"
)

type App struct {
	AdminQueries  adminQueries
	AdminCommands adminCommands
	Commands      commands
	Queries       queries
}

type adminCommands struct {
	CreateUser admincommand.CreateUserHandler
}

type adminQueries struct {
	IsAdmin adminquery.IsAdminHandler
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
	admins []int64,
	plannerDomainSerice planner.DomainService,
	usersDomainService users.DomainService,
) *App {
	return &App{
		AdminCommands: adminCommands{
			CreateUser: admincommand.NewCreateUserHandler(admins, usersDomainService),
		},
		AdminQueries: adminQueries{
			IsAdmin: adminquery.NewIsAdminHandler(admins),
		},
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
