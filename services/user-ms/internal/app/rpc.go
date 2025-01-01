package app

import (
	"commons/protos/userpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func (app *Application) RPC(srv *grpc.Server) {
	userpb.RegisterUserServiceServer(srv, app)
	grpc_health_v1.RegisterHealthServer(srv, app)
}
