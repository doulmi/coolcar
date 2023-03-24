package main

import (
	rentalpb "coolcar/rental/api/gen/v1"
	rental "coolcar/rental/internal"
	repo "coolcar/rental/repo"
	server "coolcar/shared/server"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	logger, err := server.NewZapLogger()

	if err != nil {
		log.Fatal("Cannot create logger: %v", err)
	}

	server.LoadEnv("auth/.env")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		server.GetEnvWithKey("DB_USER"),
		server.GetEnvWithKey("DB_PASSWORD"),
		server.GetEnvWithKey("DB_HOST"),
		server.GetEnvWithKey("DB_PORT"),
		server.GetEnvWithKey("DB_CONNECTION"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	repo.HandleDBMigration(db)

	logger.Sugar().Fatal(server.RunGrpcServer(&server.GrpcConfig{
		Name:              "rental",
		Addr:              ":3334",
		AuthPublicKeyFile: "shared/auth/public.key",
		Logger:            logger,
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &rental.Service{
				Logger: logger,
				DB:     db,
			})
		},
	}))
}
