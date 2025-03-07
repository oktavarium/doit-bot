package dbo

import (
	"github.com/oktavarium/doit-bot/internal/server/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GroupType int64

const (
	WithChat = iota
	WithoutChat
)

type Group struct {
	Id      primitive.ObjectID `bson:"id,omitempty"`
	OwnerId primitive.ObjectID `bson:"owner_id,omitempty"`
	TgId    int64              `bson:"tg_id,omitempty"`
	Type    GroupType          `bson:"type,omitempty"`
	Name    string             `bson:"name,omitempty"`
}

func (group Group) ToDTOGroup() *dto.Group {
	return &dto.Group{
		Id:      group.Id.Hex(),
		OwnerId: group.OwnerId.Hex(),
		TgId:    group.TgId,
		Type:    dto.GroupType(group.Type),
		Name:    group.Name,
	}
}
