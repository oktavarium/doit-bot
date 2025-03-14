package dbo

import (
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	_id         primitive.ObjectID `bson:"_id,omitempty"`
	id          string             `bson:"id,omitempty"`
	ownerId     string             `bson:"owner_id,omitempty"`
	name        string             `bson:"name,omitempty"`
	description string             `bson:"description,omitempty"`
	status      bool               `bson:"done,omitempty"`
}

func (t Task) DbId() primitive.ObjectID {
	return t._id
}

func (t Task) Id() string {
	return t.id
}

func (t Task) OwnerId() string {
	return t.ownerId
}

func (t Task) Name() string {
	return t.name
}

func (t Task) Description() string {
	return t.description
}

func (t Task) Status() bool {
	return t.status
}

func FromDomainTask(dt *planner.Task) Task {
	return Task{
		id:          dt.Id(),
		ownerId:     dt.OwnerId(),
		name:        dt.Name(),
		description: dt.Description(),
		status:      dt.Status(),
	}
}

func (t Task) ToDomainTask() (*planner.Task, error) {
	return planner.RestoreTaskFromDB(
		t.id,
		t.ownerId,
		t.name,
		t.description,
		t.status,
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
