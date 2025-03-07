package dbo

import (
	"github.com/oktavarium/doit-bot/internal/server/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	OwnerId     primitive.ObjectID `bson:"owner_id,omitempty"`
	AssigneeId  primitive.ObjectID `bson:"assignee_id,omitempty"`
	ListId      primitive.ObjectID `bson:"list_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Done        bool               `bson:"done,omitempty"`
}

func (task Task) ToDTOTask() *dto.Task {
	return &dto.Task{
		Id:          task.Id.Hex(),
		OwnerId:     task.OwnerId.Hex(),
		AssigneeId:  task.AssigneeId.Hex(),
		ListId:      task.ListId.Hex(),
		Name:        task.Name,
		Description: task.Description,
		Done:        task.Done,
	}
}
