package plants

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "plants/internal/pkg/pb/plants/v1"
)

func (h *Handler) GetPlantsWithCareRulesV1(
	ctx context.Context,
	_ *api.GetPlantsWithCareRulesV1Request,
) (*api.GetPlantsWithCareRulesV1Response, error) {
	plants, err := h.storage.GetPlantsWithCareRules(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error occured")
	}
	result := make([]*api.GetPlantsWithCareRulesV1Response_Plant, 0, len(plants))

	for _, p := range plants {
		result = append(result, &api.GetPlantsWithCareRulesV1Response_Plant{
			Species: p.Species,
			Image:   p.Image,
			Id:      p.ID.Hex(),
		})
	}
	return &api.GetPlantsWithCareRulesV1Response{
		Plants: result,
	}, nil
}
