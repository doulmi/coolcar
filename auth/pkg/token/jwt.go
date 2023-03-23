package token

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTTokenManager struct {
	NowFunc      func() time.Time
	privateKey   *rsa.PrivateKey
	tokenExpired time.Duration
}

func NewTokenGenerator(expiredIn time.Duration, privateKey *rsa.PrivateKey) *JWTTokenManager {
	return &JWTTokenManager{
		NowFunc:      time.Now,
		privateKey:   privateKey,
		tokenExpired: expiredIn,
	}
}

func (this *JWTTokenManager) GenerateToken(accountId string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)

	now := this.NowFunc()
	claims["exp"] = now.Add(this.tokenExpired).Unix()
	claims["sub"] = accountId
	claims["iss"] = "coolcar"
	claims["iat"] = now.Unix()

	tokenString, err := token.SignedString(this.privateKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
