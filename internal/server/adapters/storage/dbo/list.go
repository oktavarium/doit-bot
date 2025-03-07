package dbo

import (
	"github.com/oktavarium/doit-bot/internal/server/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type List struct {
	Id          primitive.ObjectID `bson:"id,omitempty"`
	OwnerId     primitive.ObjectID `bson:"owner_id,omitempty"`
	GroupId     primitive.ObjectID `bson:"group_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
}

func (list List) ToDTOList() *dto.List {
	dtoList := &dto.List{
		Id:          list.Id.Hex(),
		OwnerId:     list.OwnerId.Hex(),
		GroupId:     list.GroupId.Hex(),
		Name:        list.Name,
		Description: list.Description,
	}

	if !list.GroupId.IsZero() {
		dtoList.GroupId = list.GroupId.Hex()
	}

	return dtoList
}
