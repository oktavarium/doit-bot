package db

import (
	"github.com/oktavarium/doit-bot/internal/server/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type dbTask struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Owner    int64              `bson:"owner"`
	Assignee *int64             `bson:"assignee,omitempty"`
	Summary  *string            `bson:"summary,omitempty"`
	Done     *bool              `bson:"done,omitempty"`
}

func dbTaskToDtoTask(task dbTask) *dto.Task {
	return &dto.Task{
		Id:       task.Id.Hex(),
		Owner:    task.Owner,
		Assignee: task.Assignee,
		Summary:  summaryFromPointer(task.Summary),
		Done:     doneFromPointer(task.Done),
	}
}

func summaryFromPointer(summary *string) string {
	if summary == nil {
		return ""
	}

	return *summary
}

func doneFromPointer(done *bool) bool {
	if done == nil {
		return false
	}

	return *done
}
