package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/skrevolve/grpc-gateway/database"
)

const (
	port = ":50051"
)

func main() {
	w := zerolog.ConsoleWriter{Out: os.Stderr}
	l := zerolog.New(w).With().Timestamp().Caller().Logger()

	db, err := database.New()
	if err != nil {
		err = fmt.Errorf("failed to connect database: %w", err)
		l.Fatal().Err(err).Msg("failed to connect the database")
	}
	l.Info().Str("name", db.Dialect().GetName()).
		Str("database", db.Dialect().CurrentDatabase()).
		Msg("success connect to the database")

	err = database.AutoMigrate(db)
	if err != nil {
		l.Fatal().Err(err).Msg("failed to migrate database")
	}
}