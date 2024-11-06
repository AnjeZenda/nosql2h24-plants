package service

import (
	"context"

	api "plants/internal/pb/plantsapi/github.com/moevm/nosql2h24-plants/server/api/plantsapi"
)

func (s *Implementation) GetExampleV1(
	ctx context.Context,
	req *api.GetExampleV1Request,
) (*api.GetExampleV1Response, error) {
	return &api.GetExampleV1Response{}, nil
}
