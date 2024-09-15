package main

import (
	"flag"
	"log"
	"net"

	"github.com/en666ki/tgbot/api/gateway"
	"github.com/en666ki/tgbot/internal/bot_gateway"
	"github.com/en666ki/tgbot/internal/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	apiToken := flag.String("api_token", "", "Telegram bot token")
	flag.Parse()
	if *apiToken == "" {
		log.Fatal("API token is required. Pass it with --api_token flag.")
	}

	err := runGrpcServer()

	grpcClient, err := runGrpcClient()

	telegramBot, err := client.NewTelegramBot(*apiToken, grpcClient)
	if err != nil {
		log.Fatalf("Failed to start Telegram bot: %v", err)
	}

	telegramBot.Start()
}

func runGrpcServer() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	gatewayService := bot_gateway.NewGatewayService()
	gateway.RegisterGatewayServiceServer(grpcServer, gatewayService)
	reflection.Register(grpcServer)

	go func() {
		log.Println("Starting gRPC server on port :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	return nil
}

func runGrpcClient() (*grpc.ClientConn, error) {
	creds := insecure.NewCredentials()
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
