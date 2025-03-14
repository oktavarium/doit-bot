package command

import (
	"context"

	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
)

type CreateTask struct {
	OwnerId     string
	Name        string
	Description string
}

type createTaskHandler struct {
	domainService *planner.DomainService
}

type CreateTaskHandler CommandHandler[CreateTask]

func NewCreateTaskHandler(domainService *planner.DomainService) CreateTaskHandler {
	return createTaskHandler{
		domainService: domainService,
	}
}

func (h createTaskHandler) Handle(ctx context.Context, cmd CreateTask) error {
	return h.domainService.CreateTask(ctx, cmd.OwnerId, cmd.Name, cmd.Description)
}
