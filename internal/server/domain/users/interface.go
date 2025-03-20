package users

import "context"

type DomainService interface {
	CreateUser(
		ctx context.Context,
		tgId int64,
		chatTgId int64,
		username string,
	) error
	GetUserByTgId(
		ctx context.Context,
		tgId int64,
	) (*User, error)
}
