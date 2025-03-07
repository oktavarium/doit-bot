package dbo

import "go.mongodb.org/mongo-driver/bson/primitive"

type UTLink struct {
	Id     primitive.ObjectID `bson:"id,omitempty"`
	UserId primitive.ObjectID `bson:"user_id,omitempty"`
	TaskId primitive.ObjectID `bson:"task_id,omitempty"`
}
