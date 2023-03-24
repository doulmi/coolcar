package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	repo "coolcar/auth/repo"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type TokenGenerator interface {
	GenerateToken(accountId string) (string, error)
}

type Service struct {
	Logger         *zap.Logger
	TokenGenerator TokenGenerator
	DB             *gorm.DB
}

func (this *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	user := repo.Account{}
	if err := this.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return nil, status.Error(codes.InvalidArgument, "The user with this email doesn't exists")
	}

	isValid := checkPasswordHash(req.Password, user.Password)
	if !isValid {
		return nil, status.Error(codes.Unauthenticated, "Password is not matched the email")
	}

	token, err := this.TokenGenerator.GenerateToken("1")
	if err != nil {
		this.Logger.Error("Failed to create JWT Token", zap.Error(err))
	}

	return &authpb.LoginResponse{
		AccessToken: token,
	}, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
