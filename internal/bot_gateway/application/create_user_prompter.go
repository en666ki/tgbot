package application

import (
	"fmt"

	"github.com/en666ki/tgbot/internal/bot_gateway/domain"
)

type CreateUserPrompter struct {}

func NewCreateUserPrompter() *CreateUserPrompter {
	return &CreateUserPrompter{}
}

func (c *CreateUserPrompter) Prompt(user domain.User) string {
	mention := fmt.Sprintf("<a href='tg://user?id=%d'>%s</a>", user.ID, user.Username)
	return fmt.Sprintf("Привет, %s! Добро пожаловать в бот.", mention)
}
