package handlers

import (
	"auth-ms/internal/app/entities/dtos"
	"context"
)

type IUserService interface {
	RegisterUser(ctx context.Context, req dtos.CreateUserDTO) (*dtos.UserDTO, error)
}
