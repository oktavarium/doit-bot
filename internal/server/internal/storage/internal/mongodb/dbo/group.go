package dbo

import (
	"github.com/oktavarium/doit-bot/internal/server/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Group struct {
	Id      primitive.ObjectID `bson:"id,omitempty"`
	TgId    int64              `bson:"tg_id,omitempty"`
	OwnerId primitive.ObjectID `bson:"owner_id,omitempty"`
	Name    string             `bson:"name,omitempty"`
}

func (group Group) ToDTOGroup() *dto.Group {
	return &dto.Group{
		Id:      group.Id.Hex(),
		TgId:    group.TgId,
		OwnerId: group.OwnerId.Hex(),
		Name:    group.Name,
	}
}
