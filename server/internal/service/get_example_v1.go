package service

import (
	"context"

	"plants/internal/pb/plantsapi/github.com/moevm/nosql2h24-plants/server/api/plantsapi"
)

func (s *Implementation) GetExampleV1(
	ctx context.Context,
	req *plantsapi.GetExampleV1Request,
) (*plantsapi.GetExampleV1Response, error) {
	return &plantsapi.GetExampleV1Response{}, nil
}
