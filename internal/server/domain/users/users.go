package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	id       string
	tgId     int64
	chatTgId int64
	username string
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

type DomainService struct {
	repo UsersRepository
}

func NewDomainService(repo UsersRepository) *DomainService {
	return &DomainService{repo}
}

func (s *DomainService) CreateUser(
	ctx context.Context,
	tgId int64,
	chatTgId int64,
	username string,
) error {
	if err := validateTgId(tgId); err != nil {
		return fmt.Errorf("validate user tg id: %w", err)
	}

	if err := validateTgId(chatTgId); err != nil {
		return fmt.Errorf("validate chat tg id: %w", err)
	}

	if err := validatUsername(username); err != nil {
		return fmt.Errorf("validate username: %w", err)
	}

	_, err := s.repo.GetUserByTgId(ctx, tgId)
	if err != nil && !errors.Is(err, ErrUserExists) {
		return fmt.Errorf("get user by tg id: %w", err)
	}

	if errors.Is(err, ErrUserExists) {
		return err
	}

	newId, err := uuid.NewV7()
	if err != nil {
		return fmt.Errorf("generate id: %w", err)
	}

	newUser := &User{
		id:       newId.String(),
		tgId:     tgId,
		chatTgId: chatTgId,
		username: username,
	}

	if err := s.repo.CreateUser(ctx, newUser); err != nil {
		return fmt.Errorf("create new user: %w", err)
	}

	return nil
}

func (s *DomainService) GetUserByTgId(
	ctx context.Context,
	tgId int64,
) (*User, error) {
	if err := validateTgId(tgId); err != nil {
		return nil, fmt.Errorf("validate user tg id: %w", err)
	}

	user, err := s.repo.GetUserByTgId(ctx, tgId)
	if err != nil && !errors.Is(err, ErrUserExists) {
		return nil, fmt.Errorf("get user by tg id: %w", err)
	}

	return user, nil
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
	}, nil
}
