package main

import (
	rentalpb "coolcar/rental/api/gen/v1"
	rental "coolcar/rental/internal"
	server "coolcar/shared/server"
	"log"

	"google.golang.org/grpc"
)

func main() {
	logger, err := server.NewZapLogger()

	if err != nil {
		log.Fatal("Cannot create logger: %v", err)
	}

	logger.Sugar().Fatal(server.RunGrpcServer(&server.GrpcConfig{
		Name:              "rental",
		Addr:              ":3334",
		AuthPublicKeyFile: "shared/auth/public.key",
		Logger:            logger,
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &rental.Service{Logger: logger})
		},
	}))
}
