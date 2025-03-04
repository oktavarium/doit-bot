package dbo

import (
	"github.com/oktavarium/doit-bot/internal/server/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Group struct {
	Id    primitive.ObjectID   `bson:"id,omitempty"`
	TgId  int64                `bson:"tg_id,omitempty"`
	Users []primitive.ObjectID `bson:"users,omitempty"`
	Name  string               `bson:"name,omitempty"`
}

func (group Group) ToDTOGroup() *dto.Group {
	return &dto.Group{
		Id:    group.Id.Hex(),
		TgId:  group.TgId,
		Users: objectIDsToString(group.Users),
		Name:  group.Name,
	}
}
