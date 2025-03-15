package planner

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type domainService struct {
	repo PlannerRepository
}

func NewDomainService(repo PlannerRepository) DomainService {
	return &domainService{repo: repo}
}

func (s *domainService) CreateTask(
	ctx context.Context,
	ownerId string,
	name string,
	description string,
) (string, error) {
	if err := validateOwnerId(ownerId); err != nil {
		return "", fmt.Errorf("validate owner id: %w", err)
	}

	if err := validateName(name); err != nil {
		return "", fmt.Errorf("validate name: %w", err)
	}

	if err := validateDescription(name); err != nil {
		return "", fmt.Errorf("validate description: %w", err)
	}

	newId, err := uuid.NewV7()
	if err != nil {
		return "", fmt.Errorf("generate id: %w", err)
	}

	newTask := &Task{
		id:          newId.String(),
		ownerId:     ownerId,
		name:        name,
		description: description,
		_valid:      true,
	}

	if err := s.repo.CreateTask(ctx, newTask); err != nil {
		return "", fmt.Errorf("create new task: %w", err)
	}

	return newTask.id, nil
}

func (s *domainService) GetTasks(ctx context.Context, actorId string) ([]*Task, error) {
	if err := validateId(actorId); err != nil {
		return nil, fmt.Errorf("validate actor id: %w", err)
	}

	return s.repo.GetTasks(ctx, actorId)
}

func (s *domainService) GetTask(ctx context.Context, actorId string, taskId string) (*Task, error) {
	if err := validateId(taskId); err != nil {
		return nil, fmt.Errorf("validate task id: %w", err)
	}

	if err := validateId(actorId); err != nil {
		return nil, fmt.Errorf("validate actor id: %w", err)
	}

	return s.repo.GetTask(ctx, actorId, taskId)
}

func (s *domainService) DeleteTask(ctx context.Context, actorId string, taskId string) error {
	if err := validateId(taskId); err != nil {
		return fmt.Errorf("validate task id: %w", err)
	}

	if err := validateId(actorId); err != nil {
		return fmt.Errorf("validate actor id: %w", err)
	}

	return s.repo.DeleteTask(ctx, actorId, taskId)
}

func (s *domainService) UpdateTask(ctx context.Context, actorId string, task *Task) error {
	if err := validateId(actorId); err != nil {
		return fmt.Errorf("validate actor id: %w", err)
	}

	if err := isTaskValid(task); err != nil {
		return fmt.Errorf("validate task :%w", err)
	}

	return s.repo.UpdateTask(ctx, actorId, task)
}
