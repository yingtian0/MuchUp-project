package main

import (
	"MuchUp/app/config"
	grpc_controller "MuchUp/app/internal/controllers/grpc/v1"
	rest_controller "MuchUp/app/internal/controllers/http/v1"
	"MuchUp/app/internal/infrastructure/database"
	"MuchUp/app/internal/infrastructure/database/repositories"
	llmhandler "MuchUp/app/internal/infrastructure/llm"
	redisstore "MuchUp/app/internal/infrastructure/redis"
	group_service "MuchUp/app/internal/usecase/group"
	message_service "MuchUp/app/internal/usecase/message"
	user_service "MuchUp/app/internal/usecase/user"
	"MuchUp/app/pkg/logger"
	"MuchUp/app/pkg/middleware"

	"MuchUp/app/internal/infrastructure/server"

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
	config := config.LoadConfig()
	appLogger := logger.NewLogger()
	appLogger.Info("loading conifg")
	appLogger.Infof("config loaded http_port=%s grpc_port=%s db_host=%s db_name=%s", config.HTTPPort, config.GRPCPort, config.DBHost, config.DBName)

	appLogger.Info("database connecting")
	db, err := database.Connect(config)
	if err != nil {
		appLogger.Fatal("Failed to connect to database", err)
	}
	appLogger.Info("database connected")
	appLogger.Info("running database migration")
	if err = database.InitDB(db); err != nil {
		appLogger.WithError(err).Fatal("database migration failed")
	}
	appLogger.Info("Database migration completed")
	userRepo := repositories.NewUserRepository(db)
	messageRepo := repositories.NewMessageRepository(db)
	groupRepo := repositories.NewChatGroupRepository(db)
	redisClient := goredis.NewClient(&goredis.Options{
		Addr: config.GetRedisAddr(),
	})
	messageStreamStore := redisstore.NewMessageStreamStore(redisClient, 1000)
	llmConn, err := grpc.NewClient(config.GetLLMAddr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		appLogger.Fatal("Failed to connect to llm service", err)
	}
	defer llmConn.Close()
	llmHandler := llmhandler.NewHandler(llmConn, messageStreamStore)
	groupUsecase := group_service.NewGroupUsecase(groupRepo, llmHandler, appLogger)
	userUsecase := user_service.NewUserUsecase(userRepo, groupRepo, groupUsecase)
	messageUsecase := message_service.NewMessageUsecase(messageRepo, userRepo, messageStreamStore)
	RestHandler := rest_controller.NewHandler(userUsecase, messageUsecase, appLogger)

	grpcHandler := grpc_controller.NewGrpcHandler(userUsecase, messageUsecase, groupRepo, appLogger)

	HTTPRouter := RestHandler.SetupRouter()
	HTTPRouter.Use(middleware.RequestMetrics(appLogger))
	go server.StartGRPCServer(config, appLogger, grpcHandler)
	server.StartHTTPServer(config, appLogger, HTTPRouter)

}
