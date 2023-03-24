package server

import (
	"coolcar/shared/auth"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GrpcConfig struct {
	Logger            *zap.Logger
	Name              string
	Addr              string
	AuthPublicKeyFile string
	RegisterFunc      func(*grpc.Server)
}

func RunGrpcServer(c *GrpcConfig) error {
	listener, err := net.Listen("tcp", c.Addr)

	if err != nil {
		c.Logger.Fatal("Cannot listen", zap.Error(err))
	}

	var opts []grpc.ServerOption
	if c.AuthPublicKeyFile != "" {
		interceptor, err := auth.Interceptor("shared/auth/public.key")
		if err != nil {
			c.Logger.Fatal("cannot create auth interceptor", zap.Error(err))
		}
		opts = append(opts, grpc.UnaryInterceptor(interceptor))
	}

	server := grpc.NewServer(opts...)
	c.RegisterFunc(server)

	c.Logger.Info("Server started", zap.String("name", c.Name), zap.String("addr", c.Addr))
	return server.Serve(listener)
}
