package app

import (
	"commons/protos/userpb"

	"google.golang.org/grpc"
)

func (app *Application) RPC(srv *grpc.Server) {
	userpb.RegisterUserServiceServer(srv, app)
}
