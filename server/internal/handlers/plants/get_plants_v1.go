package plants

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	api "plants/internal/pkg/pb/plants/v1"
)

func (h *Handler) GetPlantsV1(
	ctx context.Context,
	req *api.GetPlantsV1Request,
) (*api.GetPlantsV1Response, error) {
	plants, err := h.storage.GetPlants(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not get plants")
	}
	result := make([]*api.GetPlantsV1Response_Plant, 0, len(plants))
	for i, p := range plants {
		result[i] = &api.GetPlantsV1Response_Plant{
			Image:     p.Image,
			Species:   p.Species,
			Price:     p.Price,
			CreatedAt: timestamppb.New(p.CreatedAt),
			Place:     p.Place,
		}
	}
	return &api.GetPlantsV1Response{
		Plants: result,
	}, nil
}
