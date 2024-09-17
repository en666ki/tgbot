package bot_gateway

import (
	"context"

	"github.com/en666ki/tgbot/api/gateway"
	"github.com/en666ki/tgbot/internal/bot_gateway/application"
	"github.com/en666ki/tgbot/internal/bot_gateway/domain"
)

type Starter interface {
	Handle(cmd domain.CreateUserCommand) error
}

type StratPrompter interface {
	Prompt(user domain.User) string
}

type GatewayService struct {
	gateway.GatewayServiceServer

	starter       Starter
	startPrompter StratPrompter
}

type DB struct {
}

func (d *DB) CreateUser(user domain.User) (id int64, err error) {
	return 0, nil
}

func (d *DB) ReadUser(id int64) (domain.User, error) {
	return domain.User{}, nil
}

func NewGatewayService() *GatewayService {
	return &GatewayService{
		starter:       application.NewCreateUserHandler(&DB{}),
		startPrompter: application.NewCreateUserPrompter(),
	}
}

func (s *GatewayService) Start(ctx context.Context, req *gateway.StartRequest) (*gateway.StartResponse, error) {
	err := s.starter.Handle(
		domain.CreateUserCommand{UserID: req.UserId, Username: req.UserName, UserFname: req.UserFname, UserSname: req.UserSname})
	if err != nil {
		return nil, err
	}

	user := domain.User{
		ID:       req.UserId,
		Username: req.UserName,
		Fname:    req.UserFname,
		Sname:    req.UserSname,
	}

	prompt := s.startPrompter.Prompt(user)

	return &gateway.StartResponse{Message: prompt}, nil
}
