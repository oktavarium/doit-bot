package dbo

import (
	"github.com/oktavarium/doit-bot/internal/server/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	TgId      int64              `bson:"tg_id,omitempty"`
	ChatTgId  int64              `bson:"chat_tg_id,omitempty"`
	FirstName string             `bson:"first_name,omitempty"`
	LastName  string             `bson:"last_name,omitempty"`
	Username  string             `bson:"username,omitempty"`
}

func (user User) ToDTOUser() *dto.User {
	return &dto.User{
		Id:   user.Id.Hex(),
		TgId: user.TgId,
	}
}
