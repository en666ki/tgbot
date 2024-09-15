package bot

import (
	"context"
	"log"

	"github.com/en666ki/tgbot/api/gateway"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/grpc"
)

type TelegramBot struct {
	bot        *tgbotapi.BotAPI
	grpcClient gateway.GatewayServiceClient
}

func NewTelegramBot(token string, grpcConn *grpc.ClientConn) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	grpcClient := gateway.NewGatewayServiceClient(grpcConn)

	return &TelegramBot{
		bot:        bot,
		grpcClient: grpcClient,
	}, nil
}

func (tb *TelegramBot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tb.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				tb.handleStartCommand(update.Message)
			}
		}
	}
}

func (tb *TelegramBot) handleStartCommand(msg *tgbotapi.Message) {
	userID := msg.From.UserName
	req := &gateway.StartRequest{
		UserId: userID,
	}

	// Отправляем gRPC запрос
	resp, err := tb.grpcClient.Start(context.Background(), req)
	if err != nil {
		log.Printf("Error calling gRPC Start method: %v", err)
		tb.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Ошибка регистрации. Попробуйте позже."))
		return
	}

	// Ответ пользователю
	tb.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, resp.Message))
}
