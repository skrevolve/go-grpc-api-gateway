package handler

import (
	"context"
	"fmt"

	"github.com/skrevolve/grpc-gateway/auth"
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

	 accessToken, err := auth.GenerateAccessToken(user.UserInfoId, user.Email)
	 if err != nil {
		msg := "internal server error"
		err := fmt.Errorf("failed to create token. %w", err)
		h.zeroLogger.Error().Err(err).Msg(msg)
		return nil, status.Error(codes.Aborted, msg)
	 }

	 return &pb.LoginResponse{AccessToken: accessToken}, nil
}

func (h *Handler) GetProfile(ctx context.Context) (*pb.GetProfileResponse, error) {
	h.zeroLogger.Info().Msg("get user profile")

	userInfo, err := auth.GetUserInfo(ctx)
	if err != nil {
		msg := "unauthenticated"
		h.zeroLogger.Error().Err(err).Msg(msg)
		return nil, status.Errorf(codes.Unauthenticated, msg)
	}

	userProfileInfo, err := h.userStore.GetProfileById(userInfo.UserInfoId)
	if err != nil {
		msg := "user not found"
		err = fmt.Errorf("token is valid but user not found: %w", err)
		h.zeroLogger.Error().Err(err).Msg(msg)
		return nil, status.Error(codes.NotFound, msg)
	}

	return &pb.GetProfileResponse{
		ImgPath: userProfileInfo.ImgPath,
		Name: userProfileInfo.Name,
		Gender: userProfileInfo.Gender,
		Email: userProfileInfo.Email,
		Lang: userProfileInfo.Lang,
	}, nil
}