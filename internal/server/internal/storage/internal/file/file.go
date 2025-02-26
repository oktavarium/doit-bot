package file

import (
	"context"
	"errors"
	"strconv"
	"sync"

	"github.com/oktavarium/doit-bot/internal/server/internal/dto"
)

type storage struct {
	tasks map[string][]*dto.Task
	mx    sync.RWMutex
}

func NewStorage() *storage {
	return &storage{
		tasks: make(map[string][]*dto.Task, 0),
	}
}

func (s *storage) CreateTask(ctx context.Context, task *dto.Task) (id string, err error) {
	if task == nil {
		return "", errors.New("empty task")
	}

	if task.Owner == nil {
		return "", errors.New("empty owner")
	}

	ownerTasks, err := s.GetTasks(ctx, *task.Owner)
	if err != nil {
		return "", errors.New("get tasks")
	}

	taskID := strconv.Itoa(len(ownerTasks))
	s.mx.Lock()
	defer s.mx.Unlock()
	task.Id = &taskID
	s.tasks[*task.Owner] = append(s.tasks[*task.Owner], task)
	return taskID, nil
}

func (s *storage) UpdateTask(ctx context.Context, task *dto.Task) error {
	return nil
}

func (s *storage) GetTask(ctx context.Context, owner, id string) (*dto.Task, error) {
	return nil, nil
}

func (s *storage) DeleteTask(ctx context.Context, owner, id string) error {
	return nil
}

func (s *storage) GetTasks(ctx context.Context, owner string) ([]*dto.Task, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	ownerTasks := s.tasks[owner]
	return ownerTasks, nil
}
