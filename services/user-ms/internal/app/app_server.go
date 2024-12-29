package app

import (
	"log"
	"net"
	"net/http"

	"user-ms/internal/pkg/protos"

	"google.golang.org/grpc"
)

func (app *Application) InitRestServer() *http.Server {
	return &http.Server{
		Addr:         app.cfg.HTTPAddr,
		Handler:      app.Routes(),
		IdleTimeout:  app.cfg.IdleTimeout,
		ReadTimeout:  app.cfg.ReadTimeout,
		WriteTimeout: app.cfg.WriteTimeout,
	}
}

func (app *Application) InitGRPCServer() (net.Listener, *grpc.Server) {
	grpcSrv := grpc.NewServer()
	protos.RegisterUserServiceServer(grpcSrv, app)

	lis, err := net.Listen("tcp", app.cfg.GRPCAddr)
	if err != nil {
		log.Fatalf("failed to init tcp listener: %v", err)
	}

	return lis, grpcSrv
}
