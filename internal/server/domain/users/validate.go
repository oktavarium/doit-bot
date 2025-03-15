package users

import (
	"github.com/google/uuid"
	"github.com/oktavarium/doit-bot/internal/doiterr"
)

func isUserValid(u *User) error {
	if u == nil {
		return ErrEmptyUser
	}

	if !u.IsValid() {
		return ErrInvalidUser
	}

	return nil
}

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
		return doiterr.WrapError(ErrInternalError, err)
	}

	return nil
}
