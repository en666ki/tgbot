package bot_gateway

import (
	"context"

	"github.com/en666ki/tgbot/api/gateway"
	"github.com/en666ki/tgbot/internal/bot_gateway/application"
	"github.com/en666ki/tgbot/internal/bot_gateway/domain"
)

type Starter interface {
	Handle(cmd domain.CreateUserCommand) (string, error)
}

type GatewayService struct {
	gateway.GatewayServiceServer

	starter Starter
}

func NewGatewayService() *GatewayService {
	return &GatewayService{starter: application.NewCreateUserHandler()}
}

func (s *GatewayService) Start(ctx context.Context, req *gateway.StartRequest) (*gateway.StartResponse, error) {
	responseMessage, err := s.starter.Handle(
		domain.CreateUserCommand{UserID: req.UserId, Username: req.UserName, UserFname: req.UserFname, UserSname: req.UserSname})
	if err != nil {
		return nil, err
	}
	return &gateway.StartResponse{Message: responseMessage}, nil
}
