package plants

import (
	"context"

	api "plants/internal/pkg/pb/plants/v1"
)

func (h *Handler) GetTradedPlantsV1(
	ctx context.Context,
	req *api.GetTradedPlantsV1Request,
) (*api.GetTradedPlantsV1Response, error) {
	return &api.GetTradedPlantsV1Response{}, nil
}
