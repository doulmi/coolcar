package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"

	"go.uber.org/zap"
)

type TokenGenerator interface {
	GenerateToken(accountId string) (string, error)
}

type Service struct {
	Logger         *zap.Logger
	TokenGenerator TokenGenerator
}

func (this *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	// if passwor and email is matched
	token, err := this.TokenGenerator.GenerateToken("1")
	if err != nil {
		this.Logger.Error("Failed to create JWT Token", zap.Error(err))
	}

	return &authpb.LoginResponse{
		AccessToken: token,
	}, nil
}
