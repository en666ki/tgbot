package application

import (
	"fmt"

	"github.com/en666ki/tgbot/internal/bot_gateway/domain"
)

type CreateUserHandler struct{}

func NewCreateUserHandler() *CreateUserHandler { return &CreateUserHandler{} }

func (h *CreateUserHandler) Handle(cmd domain.CreateUserCommand) (string, error) {
	mention := fmt.Sprintf("<a href='tg://user?id=%d'>%s</a>", cmd.UserID, cmd.UserFname)
	message := fmt.Sprintf("Привет, %s! Добро пожаловать в бот.", mention)
	return message, nil
}
