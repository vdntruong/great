package app

import (
	"context"

	"user-ms/internal/pkg/protos"
)

var _ protos.UserServiceServer = (*Application)(nil)

func (app *Application) GetByEmail(ctx context.Context, request *protos.EmailRequest) (*protos.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}
