package command

import (
	"context"
	"errors"

	"github.com/oktavarium/doit-bot/internal/doiterr"
	"github.com/oktavarium/doit-bot/internal/server/app/apperr"
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
)

type UpdateTask struct {
	ActorId     string
	TaskId      string
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
		switch {
		case errors.Is(err, planner.ErrBadId):
			return doiterr.WrapError(apperr.ErrValidationError, err)
		case errors.Is(err, planner.ErrTaskNotFound):
			return doiterr.WrapError(apperr.ErrNotFoundError, err)
		default:
			return doiterr.WrapError(apperr.ErrInternalError, err)
		}
	}

	var (
		statusChanged      bool
		nameChanged        bool
		descriptionChanged bool
	)

	if cmd.Status != nil {
		if err := task.SetStatus(cmd.ActorId, *cmd.Status); err != nil {
			switch {
			case errors.Is(err, planner.ErrForbidden):
				return doiterr.WrapError(apperr.ErrForbidden, err)
			case errors.Is(err, planner.ErrNothingChaned):
				statusChanged = false
			default:
				return doiterr.WrapError(apperr.ErrInternalError, err)
			}
		} else {
			statusChanged = true
		}
	}

	if cmd.Name != nil {
		if err := task.SetName(cmd.ActorId, *cmd.Name); err != nil {
			switch {
			case errors.Is(err, planner.ErrBadId):
				return doiterr.WrapError(apperr.ErrValidationError, err)
			case errors.Is(err, planner.ErrForbidden):
				return doiterr.WrapError(apperr.ErrForbidden, err)
			case errors.Is(err, planner.ErrNothingChaned):
				nameChanged = false
			default:
				return doiterr.WrapError(apperr.ErrInternalError, err)
			}
		} else {
			nameChanged = true
		}
	}

	if cmd.Description != nil {
		if err := task.SetDescription(cmd.ActorId, *cmd.Description); err != nil {
			switch {
			case errors.Is(err, planner.ErrBadId):
				return doiterr.WrapError(apperr.ErrValidationError, err)
			case errors.Is(err, planner.ErrForbidden):
				return doiterr.WrapError(apperr.ErrForbidden, err)
			case errors.Is(err, planner.ErrNothingChaned):
				descriptionChanged = false
			default:
				return doiterr.WrapError(apperr.ErrInternalError, err)
			}
		} else {
			descriptionChanged = true
		}
	}

	if statusChanged || nameChanged || descriptionChanged {
		if err := h.domainService.UpdateTask(ctx, cmd.ActorId, task); err != nil {
			switch {
			case errors.Is(err, planner.ErrBadId),
				errors.Is(err, planner.ErrEmptyTask),
				errors.Is(err, planner.ErrInvalidTask):
				return doiterr.WrapError(apperr.ErrValidationError, err)
			case errors.Is(err, planner.ErrTaskNotFound):
				return doiterr.WrapError(apperr.ErrNotFoundError, err)
			default:
				return doiterr.WrapError(apperr.ErrInternalError, err)
			}
		}
	}

	return nil
}
