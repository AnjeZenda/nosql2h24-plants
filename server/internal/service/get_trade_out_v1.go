package service

import (
	"context"
	api "plants/internal/pb/plantsapi/github.com/moevm/nosql2h24-plants/server/api/plantsapi"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Implementation) GetTradeOutV1(
	ctx context.Context,
	req *api.GetTradeOutV1Request,
) (*api.GetTradeOutV1Response, error) {
	user, err := s.storage.GetTrade(ctx, req.id, 1)
	if err != nil {
		return nil, status.Error(codes.Internal, "cant get trade")
	}
	var sb_of strings.Builder
	sb_of.WriteString(trade.Offerer.Surname)
	sb_of.WriteString(trade.Offerer.Name)
	sb_of.WriteString(trade.Offerer.FatherName)
	var sb_ac strings.Builder
	sb_ac.WriteString(trade.Offerer.Surname)
	sb_ac.WriteString(trade.Offerer.Name)
	sb_ac.WriteString(trade.Offerer.FatherName)
	result := &api.Trade{
		Offerer: {Name: sb_of.String(),
			Plant: {Name: trade.Offerer.Plant.Name}},
		Accepter: {Name: sb_ac.String(),
			Plant: {Name: trade.Accepter.Plant.Name}},
	}
	return &api.GetTradeOutV1Response{
		Trade: result,
	}, nil
}
