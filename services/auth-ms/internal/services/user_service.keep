package services

import (
	"auth-ms/internal/app/entities/dtos"
	"auth-ms/internal/app/entities/models"
	"context"
	"gcommons/db/dao/core"
)

type UserServiceImpl struct {
	dao core.DAO[*models.User]
}

func NewUserService(dao core.DAO[*models.User]) *UserServiceImpl {
	return &UserServiceImpl{dao: dao}
}

func (u *UserServiceImpl) RegisterUser(ctx context.Context, req *dtos.CreateUserDTO) (*dtos.UserDTO, error) {
	return nil, nil
}
