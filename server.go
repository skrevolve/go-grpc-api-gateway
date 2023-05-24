package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/skrevolve/grpc-gateway/database"
	store "github.com/skrevolve/grpc-gateway/stores"
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

	userStore := store.NewUserStore(db)

	handlers := handler.New(&l, )
}