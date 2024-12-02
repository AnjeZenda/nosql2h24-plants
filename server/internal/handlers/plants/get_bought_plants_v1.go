package plants

import (
	"context"

	api "plants/internal/pkg/pb/plants/v1"
)

func (h *Handler) GetBoughtPlantsV1(
	ctx context.Context,
	req *api.GetBoughtPlantsV1Request,
) (*api.GetBoughtPlantsV1Response, error) {
	return &api.GetBoughtPlantsV1Response{}, nil
}
