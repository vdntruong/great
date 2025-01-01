package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"commons/protos/userpb"

	"auth-ms/internal/entities/models"
	"auth-ms/internal/pkg/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type UserProviderImpl struct {
	client userpb.UserServiceClient
}

func NewUserProviderImpl(cfg *config.Config) *UserProviderImpl {
	time.Sleep(3 * time.Second)
	conn, err := grpc.NewClient(
		cfg.UserGRPCAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("user provider: could not connect: %v", err)
	}

	client := userpb.NewUserServiceClient(conn)
	return &UserProviderImpl{
		client: client,
	}
}

func (u *UserProviderImpl) VerifyUser(ctx context.Context, email string, password string) (models.Credential, error) {
	user, err := u.client.BasicAccessAuth(ctx, &userpb.BasicAuthRequest{Email: email, Password: password})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			switch statusErr.Code() {
			case codes.NotFound:
				return models.Credential{}, errors.New(statusErr.Message())
				// Handle NotFound error
			case codes.InvalidArgument:
				return models.Credential{}, errors.New(statusErr.Message())
				// Handle InvalidArgument error
			case codes.Internal:
				return models.Credential{}, errors.New(statusErr.Message()) // 500
				// Handle Internal error
			default:
				return models.Credential{}, errors.New(statusErr.Message())
				// Handle other errors
			}
		}
		return models.Credential{}, err
	}

	return models.Credential{
		ID:       user.Id,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}
