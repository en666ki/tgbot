package application

import "github.com/en666ki/tgbot/internal/bot/domain"

type CreateUserHandler struct{}

func NewCreateUserHandler() *CreateUserHandler { return &CreateUserHandler{} }

func (h *CreateUserHandler) Handle(cmd domain.CreateUserCommand) error {
	return nil
}
