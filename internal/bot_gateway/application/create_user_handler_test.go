// internal/bot/application/command_handler_test.go
package application

import (
	"fmt"
	"testing"

	"github.com/en666ki/tgbot/internal/bot_gateway/domain"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreateUserHandlerTestSuite struct {
	suite.Suite
}

type MockDB struct {
	mock.Mock
}

func TestCreateUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CreateUserHandlerTestSuite))
}

func (s *CreateUserHandlerTestSuite) TestHandle() {
	db := new(MockDB)
	handler := NewCreateUserHandler(db)
	db.On("CreateUser", domain.User{
		ID:       1,
		Username: "test",
		Fname:    "test",
		Sname:    "test",
	}).Return(int64(1), nil)

	err := handler.Handle(domain.CreateUserCommand{
		UserID:    1,
		Username:  "test",
		UserFname: "test",
		UserSname: "test",
	})
	s.Require().NoError(err)
}

func (s *CreateUserHandlerTestSuite) TestHandleError() {
	db := new(MockDB)
	handler := NewCreateUserHandler(db)
	db.On("CreateUser", domain.User{
		ID:       1,
		Username: "test",
		Fname:    "test",
		Sname:    "test",
	}).Return(int64(0), fmt.Errorf("error"))
	err := handler.Handle(domain.CreateUserCommand{
		UserID:    1,
		Username:  "test",
		UserFname: "test",
		UserSname: "test",
	})
	s.Require().Error(err)
}

func (m *MockDB) CreateUser(user domain.User) (int64, error) {
	args := m.Called(user)
	return args.Get(0).(int64), args.Error(1)
}
