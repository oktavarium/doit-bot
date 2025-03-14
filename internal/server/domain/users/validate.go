package users

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrBadTgId       = errors.New("bad tg id")
	ErrEmptyUsername = errors.New("empty username")
	ErrUserExists    = errors.New("user already exists")
)

func validateTgId(tgId int64) error {
	if tgId <= 0 {
		return ErrBadTgId
	}

	return nil
}

func validatUsername(username string) error {
	if username == "" {
		return ErrEmptyUsername
	}

	return nil
}

func validateId(id string) error {
	if id == "" {
		return nil
	}

	_, err := uuid.Parse(id)
	if err != nil {
		// TODO: return error in production
		panic(err)
	}

	return nil
}
