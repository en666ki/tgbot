// internal/bot/application/command_handler_test.go
package application

import (
	"testing"

	"github.com/en666ki/tgbot/internal/bot_gateway/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserCommand(t *testing.T) {
	handler := NewCreateUserHandler()
	command := domain.CreateUserCommand{UserID: 123, Username: "testuser", UserFname: "user", UserSname: "test"}
	response, err := handler.Handle(command)
	assert.Equal(t, "Hello, user [@user](tg://user?id=123)", response)
	assert.NoError(t, err)
}
