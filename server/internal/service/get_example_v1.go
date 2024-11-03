package service

import (
	"context"
	"plants/internal/pb/plantsapi"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Implementation) GetExampleV1(
	ctx context.Context,
	req *plantsapi.GetExampleV1Request,
) (*plantsapi.GetExampleV1Response, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented!")
}
