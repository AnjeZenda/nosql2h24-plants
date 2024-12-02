package data

import (
	"context"

	api "plants/internal/pkg/pb/data/v1"
)

func (h *Handler) ExportDBV1(
	ctx context.Context,
	req *api.ExportDBV1Request,
) (*api.ExportDBV1Response, error) {

	return &api.ExportDBV1Response{}, nil
}
