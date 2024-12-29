package repository

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	"auth-ms/internal/app/entities/models"
	"auth-ms/internal/pkg/config"
	"auth-ms/internal/pkg/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserProviderImpl struct {
	client protos.UserServiceClient
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

	client := protos.NewUserServiceClient(conn)
	return &UserProviderImpl{
		client: client,
	}
}

func (u *UserProviderImpl) VerifyUser(ctx context.Context, email string, password string) (models.Credential, error) {
	user, err := u.client.BasicAccessAuth(ctx, &protos.BasicAuthRequest{Email: email, Password: password})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			switch statusErr.Code() {
			case codes.NotFound:
				// Handle NotFound error
			case codes.InvalidArgument:
				// Handle InvalidArgument error
			case codes.Internal:
				// Handle Internal error
			default:
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
