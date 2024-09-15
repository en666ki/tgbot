package bot

import (
	"context"

	"github.com/en666ki/tgbot/api/gateway"
)

type GatewayService struct {
	gateway.UnimplementedGatewayServiceServer
}

func NewGatewayService() *GatewayService {
	return &GatewayService{}
}

func (s *GatewayService) Start(ctx context.Context, req *gateway.StartRequest) (*gateway.StartResponse, error) {
	// Логика создания пользователя
	responseMessage := "Welcome to the bot, user @" + req.UserId
	return &gateway.StartResponse{Message: responseMessage}, nil
}
