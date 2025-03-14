package dbo

import (
	"github.com/oktavarium/doit-bot/internal/server/domain/users"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	DbId     primitive.ObjectID `bson:"_id,omitempty"`
	Id       string             `bson:"id,omitempty"`
	TgId     int64              `bson:"tg_id,omitempty"`
	ChatTgId int64              `bson:"chat_tg_id,omitempty"`
	Username string             `bson:"username,omitempty"`
}

func FromDomainUser(du *users.User) User {
	return User{
		Id:       du.Id(),
		TgId:     du.TgId(),
		ChatTgId: du.ChatTgId(),
		Username: du.Username(),
	}
}

func (u User) ToDomainUser() (*users.User, error) {
	return users.RestoreUserFromDB(
		u.Id,
		u.TgId,
		u.ChatTgId,
		u.Username,
	)
}
