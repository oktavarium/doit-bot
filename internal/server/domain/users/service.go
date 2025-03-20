package users

import (
	"context"
	"errors"
	"fmt"
)

type domainService struct {
	repo UsersRepository
}

func NewDomainService(repo UsersRepository) DomainService {
	return &domainService{repo}
}

func (s *domainService) CreateUser(
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
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		return fmt.Errorf("get user by tg id: %w", err)
	}

	if err == nil {
		return ErrUserExists
	}

	newId, err := generateId()
	if err != nil {
		return fmt.Errorf("generate id: %w", err)
	}

	newUser := &User{
		id:       newId,
		tgId:     tgId,
		chatTgId: chatTgId,
		username: username,
		_valid:   true,
	}

	if err := s.repo.CreateUser(ctx, newUser); err != nil {
		return fmt.Errorf("create new user: %w", err)
	}

	return nil
}

func (s *domainService) GetUserByTgId(
	ctx context.Context,
	tgId int64,
) (*User, error) {
	if err := validateTgId(tgId); err != nil {
		return nil, fmt.Errorf("validate user tg id: %w", err)
	}

	user, err := s.repo.GetUserByTgId(ctx, tgId)
	if err != nil {
		return nil, fmt.Errorf("get user by tg id: %w", err)
	}

	return user, nil
}
