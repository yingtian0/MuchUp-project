package server

import (
	"fmt"
	"net"
	"net/http"

	"MuchUp/app/config"
	grpc_controller "MuchUp/app/internal/controllers/grpc/v1"
	"MuchUp/app/pkg/logger"
	authv1 "MuchUp/app/proto/gen/go/auth/v1"
	chatv1 "MuchUp/app/proto/gen/go/chat/v1"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func StartGRPCServer(cfg *config.Config, appLogger logger.Logger, grpcHandler *grpc_controller.GrpcHandler) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCPort))
	if err != nil {
		appLogger.Fatal("Failed to listen for gRPC", err)
	}

	s := grpc.NewServer()
	authv1.RegisterAuthServiceServer(s, grpcHandler)
	chatv1.RegisterChatServiceServer(s, grpcHandler)

	appLogger.Info("gRPC server listening at " + lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		appLogger.Fatal("Failed to serve gRPC", err)
	}
}

func StartHTTPServer(cfg *config.Config, appLogger logger.Logger, router *mux.Router) {
	httpPort := cfg.HTTPPort
	if httpPort == "" {
		httpPort = "8080"
	}

	serverAddr := fmt.Sprintf(":%s", httpPort)
	appLogger.Info("HTTP server starting on " + serverAddr)
	if err := http.ListenAndServe(serverAddr, router); err != nil {
		appLogger.Fatal("Failed to start HTTP server", err)
	}
}
