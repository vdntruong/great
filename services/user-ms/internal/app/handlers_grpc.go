package app

import (
	"context"
	"errors"

	gpassword "gcommons/password"

	"user-ms/internal/model"
	"user-ms/internal/pkg/protos"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ protos.UserServiceServer = (*Application)(nil)

func (app *Application) BasicAccessAuth(ctx context.Context, req *protos.BasicAuthRequest) (*protos.UserResponse, error) {
	user, err := app.userRepo.GetByEmail(ctx, req.GetEmail())
	if err != nil {
		return nil, err
	}

	if user == nil || errors.Is(err, model.ErrUserNotFound) {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	if !gpassword.CheckHash(req.GetPassword(), user.PasswordHash) {
		return nil, status.Error(codes.InvalidArgument, "wrong password")
	}

	res := &protos.UserResponse{
		Id:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
	return res, nil
}

func (app *Application) GetByEmail(ctx context.Context, request *protos.EmailRequest) (*protos.UserResponse, error) {
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

	res := &protos.UserResponse{
		Id:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
	return res, nil
}
