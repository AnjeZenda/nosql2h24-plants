package service

import (
	"context"
	api "plants/internal/pb/plantsapi/github.com/moevm/nosql2h24-plants/server/api/plantsapi"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Implementation) GetUserV1(
	ctx context.Context,
	req *api.GetUserV1Request,
) (*api.GetUserV1Response, error) {
	user, err := s.storage.GetUser(ctx, req.id)
	if err != nil {
		return nil, status.Error(codes.Internal, "cant get user")
	}
	result := &api.User{
		Photo:       user.Photo,
		Surname:     user.Surname,
		Name:        user.Name,
		FatherName:  user.FatherName,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
	return &api.GetUserV1Response{
		User: result,
	}, nil
}
