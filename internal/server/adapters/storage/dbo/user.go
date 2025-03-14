package dbo

import (
	"github.com/oktavarium/doit-bot/internal/server/domain/users"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	_id      primitive.ObjectID `bson:"_id,omitempty"`
	id       string             `bson:"id,omitempty"`
	tgId     int64              `bson:"tg_id,omitempty"`
	chatTgId int64              `bson:"chat_tg_id,omitempty"`
	username string             `bson:"username,omitempty"`
}

func FromDomainUser(du *users.User) User {
	return User{
		id:       du.Id(),
		tgId:     du.TgId(),
		chatTgId: du.ChatTgId(),
		username: du.Username(),
	}
}

func (u User) ToDomainUser() (*users.User, error) {
	return users.RestoreUserFromDB(
		u.id,
		u.tgId,
		u.chatTgId,
		u.username,
	)
}
