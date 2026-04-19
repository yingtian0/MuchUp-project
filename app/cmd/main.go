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
	"MuchUp/app/pkg/middleware"

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
	groupUsecase := group_service.NewGroupUsecase(groupRepo, appLogger)
	userUsecase := user_service.NewUserUsecase(userRepo, groupUsecase)
	messageUsecase := message_service.NewMessageUsecase(messageRepo, userRepo)
	RestHandler := rest_controller.NewHandler(userUsecase, messageUsecase, appLogger)

	grpcHandler := grpc_controller.NewGrpcHandler(userUsecase, messageUsecase, groupRepo, appLogger)

	HTTPRouter := RestHandler.SetupRouter()
	HTTPRouter.Use(middleware.RequestMetrics(appLogger))
	go server.StartGRPCServer(config, appLogger, grpcHandler)
	server.StartHTTPServer(config, appLogger, HTTPRouter)

}
