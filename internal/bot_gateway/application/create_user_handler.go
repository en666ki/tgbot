package application

import (
	"github.com/en666ki/tgbot/internal/bot_gateway/domain"
)

type DBUserCreator interface {
	CreateUser(user domain.User) (id int64, err error)
}

type CreateUserHandler struct {
	db DBUserCreator
}

func NewCreateUserHandler(db DBUserCreator) *CreateUserHandler { return &CreateUserHandler{db: db} }

func (h *CreateUserHandler) Handle(cmd domain.CreateUserCommand) error {
	_, err := h.db.CreateUser(domain.User{
		ID:       cmd.UserID,
		Username: cmd.Username,
		Fname:    cmd.UserFname,
		Sname:    cmd.UserSname,
	})

	if err != nil {
		return err
	}

	// structured logging
	/*
		log.WithFields(log.Fields{
			"UserID":   cmd.UserID,
			"Username": cmd.Username,
			"UserFname": cmd.UserFname,
			"UserSname": cmd.UserSname,
		}).Info("User created")
	*/
	return nil
}
