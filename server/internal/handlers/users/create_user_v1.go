// Copyright (c) 2023-2024, KNS Group LLC ("YADRO").
// All Rights Reserved.
// This software contains the intellectual property of YADRO
// or is licensed to YADRO from third parties. Use of this
// software and the intellectual property contained therein is expressly
// limited to the terms and conditions of the License Agreement under which
// it is provided by YADRO.
//

package users

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"plants/internal/models"
	api "plants/internal/pkg/pb/users/v1"
	repo "plants/internal/storage"
)

func (h *Handler) UserRegisterV1(ctx context.Context, req *api.UserRegisterV1Request) (*api.UserRegisterV1Response, error) {
	slog.Info("request", slog.String("email", req.GetEmail()), slog.String("password", req.GetPassword()))
	if req.GetEmail() == "" || req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "email or password is empty")
	}
	if _, err := h.storage.GetUserByEmail(ctx, req.GetEmail()); err != nil {
		if errors.Is(err, repo.ErrUserNotFound) {
			if err = h.storage.CreateUser(ctx, models.User{
				Surname:    req.GetSurname(),
				Name:       req.GetName(),
				FatherName: req.GetFatherName(),
				Email:      req.GetEmail(),
				Password:   req.GetPassword(),
				CreatedAt:  time.Now().UTC(),
				UpdatedAt:  time.Now().UTC(),
				Role:       models.Regular,
				Trades:     make([]primitive.ObjectID, 0),
				Plants:     make([]models.Plant, 0),
			}); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			return &api.UserRegisterV1Response{}, nil
		}
		return &api.UserRegisterV1Response{}, status.Error(codes.Internal, err.Error())
	}

	return &api.UserRegisterV1Response{}, status.Error(codes.AlreadyExists, "user already exists")
}
