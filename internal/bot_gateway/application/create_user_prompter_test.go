package application

import (
	"fmt"
	"testing"

	"github.com/en666ki/tgbot/internal/bot_gateway/domain"
	"github.com/stretchr/testify/suite"
)

type CreateUserPrompterTestSuite struct {
	suite.Suite
}

func TestCreateUserPrompterTestSuite(t *testing.T) {
	suite.Run(t, new(CreateUserPrompterTestSuite))
}

func (s *CreateUserPrompterTestSuite) TestPrompt() {
	user := domain.User{
		ID:       1,
		Username: "test",
		Fname:    "test",
		Sname:    "test",
	}
	prompter := NewCreateUserPrompter()
	prompt := prompter.Prompt(user)
	s.Require().Equal(fmt.Sprintf("Привет, <a href='tg://user?id=%d'>%s</a>! Добро пожаловать в бот.", user.ID, user.Username), prompt)
}
