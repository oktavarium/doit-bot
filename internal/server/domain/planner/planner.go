package planner

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Task struct {
	id          string
	ownerId     string
	name        string
	description string
	status      bool
}

func (t *Task) Id() string {
	return t.id
}

func (t *Task) OwnerId() string {
	return t.ownerId
}

func (t *Task) Name() string {
	return t.name
}

func (t *Task) Description() string {
	return t.description
}

func (t *Task) Status() bool {
	return t.status
}

func (t *Task) SetStatus(status bool) error {
	if t.status == status {
		return ErrNothingChaned
	}

	t.status = status
	return nil
}

type DomainService struct {
	repo PlannerRepository
}

func NewDomainService(repo PlannerRepository) *DomainService {
	return &DomainService{repo: repo}
}

func (s *DomainService) CreateTask(
	ctx context.Context,
	ownerId string,
	name string,
	description string,
) error {
	if err := validateOwnerId(ownerId); err != nil {
		return fmt.Errorf("validate owner id: %w", err)
	}

	if err := validateName(name); err != nil {
		return fmt.Errorf("validate name: %w", err)
	}

	if err := validateDescription(name); err != nil {
		return fmt.Errorf("validate description: %w", err)
	}

	newId, err := uuid.NewV7()
	if err != nil {
		return fmt.Errorf("generate id: %w", err)
	}

	newTask := &Task{
		id:          newId.String(),
		ownerId:     ownerId,
		name:        name,
		description: description,
	}

	if err := s.repo.CreateTask(ctx, newTask); err != nil {
		return fmt.Errorf("create new task: %w", err)
	}

	return nil
}

func (s *DomainService) GetTasks(ctx context.Context, actorId string) ([]*Task, error) {
	if err := validateId(actorId); err != nil {
		return nil, fmt.Errorf("validate actor id: %w", err)
	}

	return s.repo.GetTasks(ctx, actorId)
}

func (s *DomainService) GetTask(ctx context.Context, taskId string) (*Task, error) {
	if err := validateId(taskId); err != nil {
		return nil, fmt.Errorf("validate task id: %w", err)
	}

	return s.repo.GetTask(ctx, taskId)
}

func (s *DomainService) DeleteTask(ctx context.Context, actorId string, taskId string) error {
	task, err := s.GetTask(ctx, taskId)
	if err != nil {
		return fmt.Errorf("get task by id: %w", err)
	}

	if task.ownerId != actorId {
		return ErrForbidden
	}

	return s.repo.DeleteTask(ctx, taskId)
}

func (s *DomainService) UpdateTask(ctx context.Context, task *Task) error {
	oldTask, err := s.GetTask(ctx, task.id)
	if err != nil {
		return fmt.Errorf("get task by id: %w", err)
	}

	if oldTask.ownerId != task.ownerId {
		return ErrForbidden
	}

	oldTask.name = task.name
	oldTask.description = task.description
	oldTask.status = task.status

	return s.repo.UpdateTask(ctx, task)
}

func RestoreTaskFromDB(
	id string,
	ownerId string,
	name string,
	description string,
	status bool,
) (*Task, error) {
	if err := validateId(id); err != nil {
		return nil, fmt.Errorf("validate task id: %w", err)
	}

	if err := validateOwnerId(ownerId); err != nil {
		return nil, fmt.Errorf("validate owner id: %w", err)
	}

	if err := validateName(name); err != nil {
		return nil, fmt.Errorf("validate name: %w", err)
	}

	if err := validateDescription(name); err != nil {
		return nil, fmt.Errorf("validate description: %w", err)
	}

	return &Task{
		id:          id,
		ownerId:     ownerId,
		name:        name,
		description: description,
		status:      status,
	}, nil
}
