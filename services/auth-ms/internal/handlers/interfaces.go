package handlers

import (
	"auth-ms/internal/entities/dtos"
	"context"
)

type IAuthService interface {
	Login(context.Context, dtos.LoginReq) (dtos.LoginRes, error)
	VerifyAuthToken(context.Context, string) error
}
