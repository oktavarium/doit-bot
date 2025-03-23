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
	listId *string,
	name string,
	description string,
) (*Task, error) {
	return newTask(ownerId, listId, name, description)
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

func (s *domainService) GetListTasks(ctx context.Context, actorId string, listId string) ([]*Task, error) {
	if err := validateId(actorId); err != nil {
		return nil, fmt.Errorf("validate actor id: %w", err)
	}

	if err := validateId(listId); err != nil {
		return nil, fmt.Errorf("validate list id: %w", err)
	}

	return s.repo.GetListTasks(ctx, actorId, listId)
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

	return s.repo.UpdateTask(ctx, task)
}

func (s *domainService) NewList(
	ownerId string,
	name string,
	description string,
) (*List, error) {
	return newList(ownerId, name, description)
}

func (s *domainService) SaveList(
	ctx context.Context,
	list *List,
) error {
	if err := isListValid(list); err != nil {
		return fmt.Errorf("validate list :%w", err)
	}

	if err := s.repo.CreateList(ctx, list); err != nil {
		return fmt.Errorf("create new list: %w", err)
	}

	return nil
}

func (s *domainService) GetLists(ctx context.Context, actorId string) ([]*List, error) {
	if err := validateId(actorId); err != nil {
		return nil, fmt.Errorf("validate actor id: %w", err)
	}

	return s.repo.GetLists(ctx, actorId)
}

func (s *domainService) GetList(ctx context.Context, actorId string, listId string) (*List, error) {
	if err := validateId(listId); err != nil {
		return nil, fmt.Errorf("validate list id: %w", err)
	}

	if err := validateId(actorId); err != nil {
		return nil, fmt.Errorf("validate actor id: %w", err)
	}

	return s.repo.GetList(ctx, actorId, listId)
}

func (s *domainService) DeleteList(ctx context.Context, actorId string, listId string) error {
	if err := validateId(listId); err != nil {
		return fmt.Errorf("validate list id: %w", err)
	}

	if err := validateId(actorId); err != nil {
		return fmt.Errorf("validate actor id: %w", err)
	}

	return s.repo.DeleteList(ctx, actorId, listId)
}

func (s *domainService) UpdateList(ctx context.Context, actorId string, list *List) error {
	if err := validateId(actorId); err != nil {
		return fmt.Errorf("validate actor id: %w", err)
	}

	if err := isListValid(list); err != nil {
		return fmt.Errorf("validate list :%w", err)
	}

	return s.repo.UpdateList(ctx, list)
}
