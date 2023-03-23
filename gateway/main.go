package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/server"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	logger, err := server.NewZapLogger()

	if err != nil {
		log.Fatal("Cannot create logger: %v", err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseEnumNumbers: true,
				UseProtoNames:  false, // lower case with _ when true, by default is false
			}}),
	)

	serverConfigs := []struct {
		name         string
		addr         string
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{
			name:         "auth",
			addr:         ":3333",
			registerFunc: authpb.RegisterAuthServiceHandlerFromEndpoint,
		},
		{
			name:         "rental",
			addr:         ":3334",
			registerFunc: rentalpb.RegisterTripServiceHandlerFromEndpoint,
		},
	}

	for _, config := range serverConfigs {
		err := config.registerFunc(ctx, mux, config.addr, []grpc.DialOption{grpc.WithInsecure()})

		if err != nil {
			logger.Sugar().Fatal("Failed to start grpc gateway", config.name, zap.Error(err))
		}
	}

	addr := ":3000"
	logger.Sugar().Infof("GRPC Server started at %s", addr)
	logger.Sugar().Fatal(http.ListenAndServe(addr, mux))
}
