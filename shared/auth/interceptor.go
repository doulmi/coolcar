package auth

import (
	"context"
	"coolcar/shared/auth/token"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authorizationHeader = "authorization"
	bearer              = "Bearer "
)

func Interceptor(publicKeyFile string) (grpc.UnaryServerInterceptor, error) {
	file, err := os.Open(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("Cannot open public key file: %v", err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Cannot read public file: %v", err)
	}

	pk, err := jwt.ParseRSAPublicKeyFromPEM(bytes)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse rsa public key: %v", err)
	}

	i := interceptor{
		verifier: &token.JWTTokenVerifier{
			PublicKey: pk,
		},
	}
	return i.HandleRequest, nil
}

type tokenVerifier interface {
	Verify(token string) (string, error)
}

type interceptor struct {
	verifier tokenVerifier
}

func (this *interceptor) HandleRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	token, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}

	accountId, err := this.verifier.Verify(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Token not valid: %v", err)
	}

	return handler(ContextWithAccountID(ctx, accountId), req)
}

func tokenFromContext(ctx context.Context) (string, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}

	token := ""
	for _, v := range meta[authorizationHeader] {
		if strings.HasPrefix(v, bearer) {
			token = v[len(bearer):]
		}
	}

	if token == "" {
		return "", status.Error(codes.Unauthenticated, "")
	}

	if strings.Contains(token, " ") {
		return "", status.Error(codes.Unauthenticated, "")
	}
	return token, nil
}

type accountIDKey struct{}

func ContextWithAccountID(c context.Context, accountId string) context.Context {
	return context.WithValue(c, accountIDKey{}, accountId)
}

func AccountIDFromContext(c context.Context) (string, error) {
	value := c.Value(accountIDKey{})

	accountId, ok := value.(string)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}
	return accountId, nil
}
