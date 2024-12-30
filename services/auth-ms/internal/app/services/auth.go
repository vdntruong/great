package services

import (
	"context"

	"auth-ms/internal/app/entities/dtos"
	"auth-ms/internal/pkg/config"
)

type AuthServiceImpl struct {
	cfg          *config.Config
	tokenAdaptor TokenAdaptor
	userProvider IUserProvider
}

func NewAuthServiceImpl(
	cfg *config.Config,
	tokenAdaptor TokenAdaptor,
	userProvider IUserProvider,
) *AuthServiceImpl {
	return &AuthServiceImpl{
		cfg:          cfg,
		tokenAdaptor: tokenAdaptor,
		userProvider: userProvider,
	}
}

func (a *AuthServiceImpl) Login(ctx context.Context, req dtos.LoginReq) (dtos.LoginRes, error) {
	user, err := a.userProvider.VerifyUser(ctx, req.Email, req.Password)
	if err != nil {
		return dtos.LoginRes{}, err
	}

	accessToken, refreshToken, err := a.tokenAdaptor.GenerateTokenPair(user.ID, user.Username)
	if err != nil {
		return dtos.LoginRes{}, err
	}

	res := dtos.LoginRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return res, nil
}

func (a *AuthServiceImpl) VerifyAuthToken(ctx context.Context, token string) error {
	if err := a.tokenAdaptor.ValidateToken(token); err != nil {
		return err
	}
	return nil
}
