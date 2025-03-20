package planner

import (
	"context"
	"fmt"
)

type domainService struct {
	repo PlannerRepository
}

func NewDomainService(repo PlannerRepository) DomainService {
	return &domainService{repo: repo}
}

func (s *domainService) NewTask(
	ownerId string,
	name string,
	description string,
) (*Task, error) {
	return newTask(ownerId, name, description)
}

func (s *domainService) SaveTask(
	ctx context.Context,
	task *Task,
) error {
	if err := isTaskValid(task); err != nil {
		return fmt.Errorf("validate task :%w", err)
	}

	if err := s.repo.CreateTask(ctx, task); err != nil {
		return fmt.Errorf("create new task: %w", err)
	}

	return nil
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
