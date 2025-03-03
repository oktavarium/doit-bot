package db

import (
	"github.com/oktavarium/doit-bot/internal/server/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type dbTask struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Owner       primitive.ObjectID `bson:"owner,omitempty"`
	Assignee    primitive.ObjectID `bson:"assignee,omitempty"`
	List        primitive.ObjectID `bson:"list,omitempty"`
	Summary     string             `bson:"summary,omitempty"`
	Description string             `bson:"description,omitempty"`
	Done        bool               `bson:"done,omitempty"`
}

type dbUser struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	TgId      int64              `bson:"tg_id,omitempty"`
	FirstName string             `bson:"first_name,omitempty"`
	LastName  string             `bson:"last_name,omitempty"`
	Username  string             `bson:"username,omitempty"`
}

func dbUserToDTOUser(user dbUser) *dto.User {
	return &dto.User{
		Id:   user.Id.Hex(),
		TgId: user.TgId,
	}
}

func dbTaskToDtoTask(task dbTask) *dto.Task {
	return &dto.Task{
		Id:       task.Id.Hex(),
		Owner:    task.Owner.Hex(),
		Assignee: task.Assignee.Hex(),
		List:     stringToPointer(task.List.Hex()),
		Summary:  task.Summary,
		Done:     task.Done,
	}
}

func stringFromPointer(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

func stringToPointer(s string) *string {
	return &s
}

func boolFromPointer(b *bool) bool {
	if b == nil {
		return false
	}

	return *b
}
