package users

import (
	"errors"
)

var (
	ErrEmptyUser           = errors.New("empty user")
	ErrInvalidUser         = errors.New("user created outside domain layer")
	ErrBadTgId             = errors.New("bad tg id")
	ErrEmptyUsername       = errors.New("empty username")
	ErrUserExists          = errors.New("user already exists")
	ErrUserNotFound        = errors.New("user not found")
	ErrInternalError       = errors.New("internal error")
	ErrInfrastructureError = errors.New("infrastructure error")
)
