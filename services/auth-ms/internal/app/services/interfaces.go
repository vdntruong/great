package services

import (
	"auth-ms/internal/app/entities/models"
	"context"
)

type IUserProvider interface {
	VerifyUser(ctx context.Context, email string, password string) (models.Credential, error)
}
