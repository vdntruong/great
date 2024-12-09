package server

import (
	"auth-ms/internal/app/handlers"
	"auth-ms/internal/app/repository"
	"auth-ms/internal/pkg/config"
	"net"

	authpb "auth-ms/pkg/protos"
	"google.golang.org/grpc"
)

func StartGRPCServer(cfg *config.Config, userRepo repository.UserRepository) error {
	lis, err := net.Listen("tcp", cfg.GRPCAddress)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	authHandler := handlers.NewAuthGRPCHandler(userRepo)
	authpb.RegisterAuthServiceServer(grpcServer, authHandler)

	return grpcServer.Serve(lis)
}
