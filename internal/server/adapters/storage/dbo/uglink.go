package dbo

import "go.mongodb.org/mongo-driver/bson/primitive"

type UGLink struct {
	Id      primitive.ObjectID `bson:"id,omitempty"`
	UserId  primitive.ObjectID `bson:"user_id,omitempty"`
	GroupId primitive.ObjectID `bson:"group_id,omitempty"`
}
