// internal/bot/application/command_handler_test.go
package application

import (
	"testing"

	"github.com/en666ki/tgbot/internal/bot/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserCommand(t *testing.T) {
	handler := NewCreateUserHandler()
	command := domain.CreateUserCommand{UserID: "123", Username: "testuser"}
	err := handler.Handle(command)
	assert.NoError(t, err)
}
