package data

import (
	"context"

	api "plants/internal/pkg/pb/data/v1"
)

func (h *Handler) ImportDBV1(
	ctx context.Context,
	req *api.ImportDBV1Request,
) (*api.ImportDBV1Response, error) {
	err := h.storage.ImportDB(ctx, req.Db)
	if err != nil {
		return &api.ImportDBV1Response{}, err
	}
	return &api.ImportDBV1Response{}, nil
}
