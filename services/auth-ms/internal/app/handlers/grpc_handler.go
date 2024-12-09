package handlers

import (
	"auth-ms/internal/app/models"
	"auth-ms/internal/app/repository"
	authpb "auth-ms/pkg/protos"
	"context"
	"gcommons/authen"
	"gcommons/password"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthGRPCHandler struct {
	authpb.UnimplementedAuthServiceServer
	userRepo repository.UserRepository
}

func NewAuthGRPCHandler(userRepo repository.UserRepository) *AuthGRPCHandler {
	return &AuthGRPCHandler{
		userRepo: userRepo,
	}
}

func (h *AuthGRPCHandler) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &authpb.RegisterResponse{
		UserId: user.ID.String(),
	}, nil
}

func (h *AuthGRPCHandler) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	user, err := h.userRepo.FindByUsername(ctx, req.Username)
	if err != nil || !password.CheckPasswordHash(req.Password, user.Password) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid username or password")
	}

	tokenPair, err := authen.GenerateTokenPair(user.ID, user.Username)
	if err != nil {
		return nil, err
	}

	return &authpb.LoginResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}
