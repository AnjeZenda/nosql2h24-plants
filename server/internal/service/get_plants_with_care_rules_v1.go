package service

import (
	"context"
	api "plants/internal/pb/plantsapi/github.com/moevm/nosql2h24-plants/server/api/plantsapi"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Implementation) GetPlantsWithCareRulesV1(
	ctx context.Context,
	req *api.GetPlantsWithCareRulesV1Request,
) (*api.GetPlantsWithCareRulesV1Response, error) {
	plants, err := s.storage.GetPlantsWithCareRules(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error occured")
	}
	result := make([]*api.GetPlantsWithCareRulesV1Response_Plant, 0, len(plants))

	for _, p := range plants {
		result = append(result, &api.GetPlantsWithCareRulesV1Response_Plant{
			Species: p.Species,
			Image:   p.Image,
			Id:      p.ID.String(),
		})
	}
	return &api.GetPlantsWithCareRulesV1Response{
		Plants: result,
	}, nil
}
