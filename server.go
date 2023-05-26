package main

import (
	"fmt"
	"net"
	"os"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/rs/zerolog"
	"github.com/skrevolve/grpc-gateway/database"
	handler "github.com/skrevolve/grpc-gateway/handlers"
	pb "github.com/skrevolve/grpc-gateway/protos"
	"github.com/skrevolve/grpc-gateway/stores"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	w := zerolog.ConsoleWriter{Out: os.Stderr}
	zeroLogger := zerolog.New(w).With().Timestamp().Caller().Logger()

	db, err := database.New()
	if err != nil {
		err = fmt.Errorf("failed to connect database: %w", err)
		zeroLogger.Fatal().Err(err).Msg("failed to connect the database")
	}
	zeroLogger.Info().Str("name", db.Name()).
		Str("database", db.Migrator().CurrentDatabase()).
		Msg("success connect to the database")

	err = database.AutoMigrate(db)
	if err != nil {
		zeroLogger.Fatal().Err(err).Msg("failed to migrate database")
	}

	userStore := stores.NewUserStore(db)

	handlers := handler.New(&zeroLogger, userStore)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		zeroLogger.Panic().Err(fmt.Errorf("faield to listen: %w", port))
	}

	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(),
		),
	)
	pb.RegisterUsersServer(grpcServer, handlers)

}