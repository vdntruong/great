package app

import (
	"context"
	"errors"
	"google.golang.org/grpc/health/grpc_health_v1"

	gpassword "commons/password"
	"commons/protos/userpb"

	"user-ms/internal/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ userpb.UserServiceServer = (*Application)(nil)

func (app *Application) BasicAccessAuth(ctx context.Context, req *userpb.BasicAuthRequest) (*userpb.UserResponse, error) {
	user, err := app.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil || errors.Is(err, model.ErrUserNotFound) {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	if !gpassword.CheckHash(req.Password, user.PasswordHash) {
		return nil, status.Error(codes.InvalidArgument, "wrong password")
	}

	res := &userpb.UserResponse{
		Id:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
	return res, nil
}

func (app *Application) GetByEmail(ctx context.Context, request *userpb.EmailRequest) (*userpb.UserResponse, error) {
	founded, err := app.userRepo.FindByEmail(ctx, request.GetEmail())
	if err != nil {
		return nil, err
	}
	if !founded {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	user, err := app.userRepo.GetByEmail(ctx, request.GetEmail())
	if err != nil {
		return nil, err
	}

	res := &userpb.UserResponse{
		Id:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
	return res, nil
}

func (app *Application) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}
