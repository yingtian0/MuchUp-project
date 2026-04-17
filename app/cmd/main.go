package main

import (
	"MuchUp/app/config"
	group_service "MuchUp/app/internal/application/service/group"
	message_service "MuchUp/app/internal/application/service/message"
	user_service "MuchUp/app/internal/application/service/user"
	grpc_controller "MuchUp/app/internal/controllers/grpc/v1"
	rest_controller "MuchUp/app/internal/controllers/http/v1"
	"MuchUp/app/internal/infrastructure/database"
	"MuchUp/app/internal/infrastructure/database/repositories"
	"MuchUp/app/pkg/logger"
	"log"

	"MuchUp/app/internal/infrastructure/auth"
	"MuchUp/app/internal/infrastructure/server"
)

// @Title MuchUp API
// @version 1.0
// @description This is the MuchUp server
// @host localhost:8080
// @Basepath /api/v1
func main() {
	config := config.LoadConfig()
	appLogger := logger.NewLogger()

	db, err := database.Connect(config)
	if err != nil {
		appLogger.Fatal("Failed to connect to database", err)
	}

	database.InitDB(db)
	appLogger.Info("Database migration completed")
	JWTValidator := auth.NewJWTValidator(config.SecretKey, "Much-Up", "users")
	userRepo := repositories.NewUserRepository(db)
	messageRepo := repositories.NewMessageRepository(db)
	groupRepo := repositories.NewChatGroupRepository(db)
	groupUsecase := group_service.NewGroupUsecase(groupRepo, appLogger)
	userUsecase := user_service.NewUserUsecase(userRepo, groupUsecase)
	messageUsecase := message_service.NewMessageUsecase(messageRepo, userRepo)
	RestHandler := rest_controller.NewHandler(userUsecase, messageUsecase, appLogger)

	grpcHandler := grpc_controller.NewGrpcHandler(userUsecase, messageUsecase, groupRepo, appLogger)

	log.Println("start serve")
	go server.StartGRPCServer(config, appLogger, grpcHandler)
	go server.StartHTTPServer(config, appLogger, RestHandler.SetupRoutes(JWTValidator))

}
