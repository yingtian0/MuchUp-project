package main

import (
	"context"

	"MuchUp/app/config"
	grpc_controller "MuchUp/app/internal/controllers/grpc/v1"
	rest_controller "MuchUp/app/internal/controllers/http/v1"
	"MuchUp/app/internal/infrastructure/postgres"
	redisstore "MuchUp/app/internal/infrastructure/redis"
	"MuchUp/app/internal/infrastructure/server"
	message_service "MuchUp/app/internal/usecase/message"
	user_service "MuchUp/app/internal/usecase/user"
	"MuchUp/app/pkg/logger"
	"MuchUp/app/pkg/middleware"

	goredis "github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// @Title MuchUp API
// @version 1.0
// @description This is the MuchUp server
// @host localhost:8080
// @Basepath /api/v1
func main() {
	ctx := context.Background()
	config := config.LoadConfig()
	appLogger := logger.NewLogger()
	appLogger.Info("loading conifg")
	appLogger.Infof("config loaded http_port=%s grpc_port=%s db_host=%s db_name=%s", config.HTTPPort, config.GRPCPort, config.DBHost, config.DBName)

	redisClient := goredis.NewClient(&goredis.Options{
		Addr: config.GetRedisAddr(),
	})
	messageStreamStore := redisstore.NewMessageStreamStore(redisClient, 1000)

	llmConn, err := grpc.NewClient(config.GetLLMAddr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		appLogger.WithError(err).Fatal("Failed to connect to llm service")
	}

	defer llmConn.Close()

	appLogger.Info("database connecting")

	db, err := postgres.Connect(ctx, config)
	if err != nil {
		appLogger.WithError(err).Fatal("Failed to connect to database")
	}
	defer db.Close()

	appLogger.Info("database connected")

	userRepository := postgres.NewUserRepository(db)
	messageRepository := postgres.NewMessageRepository(db)

	userUsecase := user_service.NewUserUsecase(userRepository)
	messageUsecase := message_service.NewMessageUsecase(messageRepository, userRepository, messageStreamStore)
	RestHandler := rest_controller.NewHandler(userUsecase, messageUsecase, appLogger)

	grpcHandler := grpc_controller.NewGrpcHandler(userUsecase, messageUsecase, appLogger)

	HTTPRouter := RestHandler.SetupRouter()
	HTTPRouter.Use(middleware.RequestMetrics(appLogger))

	go server.StartGRPCServer(config, appLogger, grpcHandler)

	server.StartHTTPServer(config, appLogger, HTTPRouter)
}
