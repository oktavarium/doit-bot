package users

import "context"

type UsersRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByTgId(ctx context.Context, tgId int64) (*User, error)
}
