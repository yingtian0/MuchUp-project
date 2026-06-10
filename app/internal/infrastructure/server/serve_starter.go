package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"MuchUp/app/config"
	grpc_controller "MuchUp/app/internal/controllers/grpc/v1"
	"MuchUp/app/pkg/logger"
	authv1 "MuchUp/app/proto/gen/go/auth/v1"
	chatv1 "MuchUp/app/proto/gen/go/chat/v1"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	grpc_health_v1 "google.golang.org/grpc/health/grpc_health_v1"
)

func StartGRPCServer(cfg *config.Config, appLogger logger.Logger, grpcHandler *grpc_controller.GrpcHandler) {
	var listenConfig net.ListenConfig

	lis, err := listenConfig.Listen(context.Background(), "tcp", fmt.Sprintf(":%s", cfg.GRPCPort))
	if err != nil {
		appLogger.WithError(err).Fatal("Failed to listen for gRPC")
	}

	s := grpc.NewServer()
	authv1.RegisterAuthServiceServer(s, grpcHandler)
	chatv1.RegisterChatServiceServer(s, grpcHandler)

	healthServer := health.NewServer()
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus(authv1.AuthService_ServiceDesc.ServiceName, grpc_health_v1.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus(chatv1.ChatService_ServiceDesc.ServiceName, grpc_health_v1.HealthCheckResponse_SERVING)
	grpc_health_v1.RegisterHealthServer(s, healthServer)

	appLogger.Infof("gRPC server listening at %s", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		appLogger.WithError(err).Fatal("Failed to serve gRPC")
	}
}

func StartHTTPServer(cfg *config.Config, appLogger logger.Logger, router *mux.Router) {
	httpPort := cfg.HTTPPort
	if httpPort == "" {
		httpPort = "8080"
	}

	serverAddr := fmt.Sprintf(":%s", httpPort)
	appLogger.Infof("HTTP server starting on %s", serverAddr)

	server := &http.Server{
		Addr:              serverAddr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		appLogger.WithError(err).Fatal("Failed to start HTTP server")
	}
}
