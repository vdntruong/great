package services

import (
	"auth-ms/internal/app/entities/dtos"
	"context"
	"gcommons/authen"
)

type AuthServiceImpl struct {
	userProvider IUserProvider
}

func NewAuthServiceImpl(userProvider IUserProvider) *AuthServiceImpl {
	return &AuthServiceImpl{
		userProvider: userProvider,
	}
}

func (a *AuthServiceImpl) Login(ctx context.Context, req dtos.LoginReq) (dtos.LoginRes, error) {
	user, err := a.userProvider.VerifyUser(ctx, req.Email, req.Password)
	if err != nil {
		return dtos.LoginRes{}, err
	}

	tokenPair, err := authen.GenerateTokenPair(user.ID, user.Username)
	if err != nil {
		return dtos.LoginRes{}, err
	}

	res := dtos.LoginRes{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}
	return res, nil
}
