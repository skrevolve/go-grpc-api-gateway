package handler

import (
	"github.com/rs/zerolog"
	store "github.com/skrevolve/grpc-gateway/stores"
)

// Handler definition
type Handler struct {
	zeroLogger	*zerolog.Logger
	userStore		*store.UserStore
}

// New returns a new handler with logger and database
func New(zeroLogger *zerolog.Logger, userStore *store.UserStore) *Handler {
	return &Handler{zeroLogger: zeroLogger, userStore: userStore}
}
