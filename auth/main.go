package main

import (
	authpb "coolcar/auth/api/gen/v1"
	auth "coolcar/auth/internal"
	"coolcar/auth/pkg/token"
	"coolcar/shared/server"
	"crypto/rsa"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := server.NewZapLogger()

	if err != nil {
		log.Fatal("Cannot create logger: %v", err)
	}

	logger.Sugar().Fatal(server.RunGrpcServer(&server.GrpcConfig{
		Name:   "auth",
		Addr:   ":3333",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			authpb.RegisterAuthServiceServer(s, &auth.Service{Logger: logger, TokenGenerator: token.NewTokenGenerator(
				2*time.Hour,
				getPrivateKey(logger),
			)})
		},
	}))

}

func getPrivateKey(logger *zap.Logger) *rsa.PrivateKey {
	pkFile, err := os.Open("auth/private.key")
	if err != nil {
		logger.Fatal("Cannot open private key", zap.Error(err))
	}

	pkBytes, err := ioutil.ReadAll(pkFile)
	if err != nil {
		logger.Fatal("Cannot read private key", zap.Error(err))
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		logger.Fatal("Cannot parse private key", zap.Error(err))
	}

	return privateKey
}
