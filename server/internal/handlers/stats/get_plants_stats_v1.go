package stats

import (
	"context"
	api "plants/internal/pkg/pb/stats/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) GetPlantsStatsV1(ctx context.Context,
	req *api.GetPlantsStatsV1Request,
) (*api.GetPlantsStatsV1Response, error) {
	fltr := parseFilter(req.Filter)
	plants, err := h.storage.GetPlantsStats(ctx, fltr)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error occured %v", err)
	}
	stats := make([]*api.GetPlantsStatsV1Response_StatsInfo, len(plants.Info))
	for i, p := range plants.Info {
		t, err := time.Parse("2006-01-02", p.Date)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Error occured %v", err)
		}
		stats[i] = &api.GetPlantsStatsV1Response_StatsInfo{
			Date:    timestamppb.New(t),
			Count:   p.Count,
			Species: p.Species,
		}
	}
	return &api.GetPlantsStatsV1Response{
		Count: plants.Count,
		Stats: stats,
	}, nil
}
