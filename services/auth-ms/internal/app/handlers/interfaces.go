package handlers

import (
	"context"

	"auth-ms/internal/app/entities/dtos"
)

type IAuthService interface {
	Login(context.Context, dtos.LoginReq) (dtos.LoginRes, error)
	VerifyAuthToken(context.Context, string) error
}
