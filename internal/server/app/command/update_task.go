package command

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/server/app/apperr"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
)

type UpdateTask struct {
	ActorId     string
	TaskId      string
	ListId      *string
	Status      *bool
	Name        *string
	Description *string
}

type updateTaskHandler struct {
	domainService planner.DomainService
}

type UpdateTaskHandler CommandHandler[UpdateTask]

func NewUpdateTaskHandler(domainService planner.DomainService) UpdateTaskHandler {
	return updateTaskHandler{
		domainService: domainService,
	}
}

func (h updateTaskHandler) Handle(ctx context.Context, cmd UpdateTask) error {
	task, err := h.domainService.GetTask(ctx, cmd.ActorId, cmd.TaskId)
	if err != nil {
		return apperr.FromPlannerError(err)
	}

	var (
		statusChanged      bool
		nameChanged        bool
		descriptionChanged bool
		listChanged        bool
	)

	if cmd.Status != nil {
		if err := task.SetStatus(cmd.ActorId, *cmd.Status); err != nil {
			if !errors.Is(err, planner.ErrNothingChaned) {
				return apperr.FromPlannerError(err)
			}
		} else {
			statusChanged = true
		}
	}

	if cmd.ListId != nil {
		if err := task.SetListId(cmd.ActorId, *cmd.ListId); err != nil {
			if !errors.Is(err, planner.ErrNothingChaned) {
				return apperr.FromPlannerError(err)
			}
		} else {
			listChanged = true
		}
	}

	if cmd.Name != nil {
		if err := task.SetName(cmd.ActorId, *cmd.Name); err != nil {
			if !errors.Is(err, planner.ErrNothingChaned) {
				return apperr.FromPlannerError(err)
			}
		} else {
			nameChanged = true
		}
	}

	if cmd.Description != nil {
		if err := task.SetDescription(cmd.ActorId, *cmd.Description); err != nil {
			if !errors.Is(err, planner.ErrNothingChaned) {
				return apperr.FromPlannerError(err)
			}
		} else {
			descriptionChanged = true
		}
	}

	if statusChanged || listChanged || nameChanged || descriptionChanged {
		if err := h.domainService.UpdateTask(ctx, cmd.ActorId, task); err != nil {
			return apperr.FromPlannerError(err)
		}
	}

	return nil
}
