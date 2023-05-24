package handler

import (
	"context"
	"fmt"

	pb "github.com/skrevolve/grpc-gateway/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	h.zeroLogger.Info().Interface("req", req).Msg("login user")

	user ,err := h.userStore.GetByEmailAndPassword(req.GetEmail(), req.GetPassword())
	if err != nil {
		msg := "invalid email or password"
		err = fmt.Errorf("failed to login due to wrong email: %w", err)
		h.zeroLogger.Error().Err(err).Msg(msg)
		return nil, status.Error(codes.InvalidArgument, msg)
	}

	 if !user.CheckPassword(req.GetPassword()) {
		h.zeroLogger.Error().Msgf("failed to login due to receive wrong password: %s", user.Email)
		return nil, status.Error(codes.InvalidArgument, "invalid email or password")
	 }

	 accessToken, err :=
}