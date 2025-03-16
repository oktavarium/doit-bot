package users

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/oktavarium/doit-bot/internal/doiterr"
)

type User struct {
	id       string
	tgId     int64
	chatTgId int64
	username string
	_valid   bool
}

func (u *User) Id() string {
	return u.id
}

func (u *User) TgId() int64 {
	return u.tgId
}

func (u *User) ChatTgId() int64 {
	return u.chatTgId
}

func (u *User) Username() string {
	return u.username
}

func (u *User) IsValid() bool {
	return u._valid
}

func generateId() (string, error) {
	newId, err := uuid.NewV7()
	if err != nil {
		return "", doiterr.WrapError(ErrInternalError, err)
	}

	return newId.String(), nil
}

func RestoreUserFromDB(
	id string,
	tgId int64,
	chatTgId int64,
	username string,
) (*User, error) {
	if err := validateId(id); err != nil {
		return nil, fmt.Errorf("validate task id: %w", err)
	}

	if err := validateTgId(tgId); err != nil {
		return nil, fmt.Errorf("validate user tg id: %w", err)
	}

	if err := validateTgId(chatTgId); err != nil {
		return nil, fmt.Errorf("validate chat tg id: %w", err)
	}

	if err := validatUsername(username); err != nil {
		return nil, fmt.Errorf("validate username: %w", err)
	}

	return &User{
		id:       id,
		tgId:     tgId,
		chatTgId: chatTgId,
		username: username,
		_valid:   true,
	}, nil
}
