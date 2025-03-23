package dbo

import (
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	DbId        primitive.ObjectID `bson:"_id,omitempty"`
	Id          string             `bson:"id,omitempty"`
	OwnerId     string             `bson:"owner_id,omitempty"`
	ListId      string             `bson:"list_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Status      bool               `bson:"status,omitempty"`
}

func FromDomainTask(dt *planner.Task) Task {
	return Task{
		Id:          dt.Id(),
		OwnerId:     dt.OwnerId(),
		ListId:      dt.ListId(),
		Name:        dt.Name(),
		Description: dt.Description(),
		Status:      dt.Status(),
	}
}

func (t Task) ToDomainTask() (*planner.Task, error) {
	return planner.RestoreTaskFromDB(
		t.Id,
		t.OwnerId,
		t.ListId,
		t.Name,
		t.Description,
		t.Status,
	)
}

func TasksToDomainTasks(tasks []Task) ([]*planner.Task, error) {
	if tasks == nil {
		return nil, nil
	}
	result := make([]*planner.Task, 0, len(tasks))
	for _, t := range tasks {
		domainTask, err := t.ToDomainTask()
		if err != nil {
			return nil, err
		}
		result = append(result, domainTask)
	}
	return result, nil
}
