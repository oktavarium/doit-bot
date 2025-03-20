package users

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
	ctrl                *gomock.Controller
	mockUsersRepository *MockUsersRepository
}

func (s *ServiceSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mockUsersRepository = NewMockUsersRepository(s.ctrl)
}

// Очистка после каждого теста
func (s *ServiceSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) TestCreateUser() {
	s.T().Parallel()

	ctx := context.Background()
	var tgId int64 = 10
	var chatTgId int64 = 10
	username := "user"

	gomock.InOrder(
		s.mockUsersRepository.EXPECT().GetUserByTgId(ctx, tgId).Return(nil, ErrUserNotFound),
		s.mockUsersRepository.EXPECT().CreateUser(ctx, gomock.Any()).Return(nil),
	)

	usersService := NewDomainService(s.mockUsersRepository)

	err := usersService.CreateUser(ctx, tgId, chatTgId, username)
	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) TestGetUserByTgId() {
	s.T().Parallel()

	ctx := context.Background()
	var tgId int64 = 10

	gomock.InOrder(
		s.mockUsersRepository.EXPECT().GetUserByTgId(ctx, tgId).Return(nil, ErrUserNotFound),
	)

	usersService := NewDomainService(s.mockUsersRepository)

	user, err := usersService.GetUserByTgId(ctx, tgId)
	assert.Nil(s.T(), user)
	assert.ErrorIs(s.T(), err, ErrUserNotFound)
}
