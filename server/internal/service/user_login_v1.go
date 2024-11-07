package service

import (
	"context"
	api "plants/internal/pb/plantsapi/github.com/moevm/nosql2h24-plants/server/api/plantsapi"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Implementation) UserLoginV1(
	ctx context.Context,
	req *api.UserLoginV1Request,
) (*api.UserLoginV1Response, error) {
	user, err := s.storage.SearchUser(ctx, req.Login, req.Password)
	if err != nil {
		return nil, status.Error(codes.NotFound, "could not find user")
	}
	return &api.UserLoginV1Response{
		id: user,
	}, nil
}
