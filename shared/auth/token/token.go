package token

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

func (this *JWTTokenVerifier) Verify(token string) (string, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return "", fmt.Errorf("Cannot parse token: %v", token)
		}
		return this.PublicKey, nil
	})

	ti, _ := t.Claims.GetExpirationTime()
	ttt, _ := time.Parse(time.DateTime, "2023-03-23 14:08:53")
	fmt.Print(ttt.Unix())

	fmt.Printf("%v", ti.Time)

	if err != nil {
		return "", fmt.Errorf("Cannot parse token: %v", err)
	}

	if !t.Valid {
		return "", fmt.Errorf("Not a valid token")
	}

	subject, err := t.Claims.GetSubject()
	if err != nil {
		return "", fmt.Errorf("Cannot get subject")
	}

	return subject, nil
}
