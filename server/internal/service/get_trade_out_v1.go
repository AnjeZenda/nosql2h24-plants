package service

import (
	"context"
	api "plants/internal/pb/plantsapi/github.com/moevm/nosql2h24-plants/server/api/plantsapi"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Implementation) GetTradeOutV1(
	ctx context.Context,
	req *api.GetTradeOutV1Request,
) (*api.GetTradeOutV1Response, error) {
	trade, err := s.storage.GetTrade(ctx, req.id, 1)
	if err != nil {
		return nil, status.Error(codes.Internal, "cant get trade")
	}
	result := &api.Trade{
		Offerer: &api.TradeUser{Surname: trade.Offerer.Surname,
			Name:       trade.Offerer.Name,
			FatherName: trade.Offerer.FatherName,
			Plant:      &api.TradePlant{Name: trade.Offerer.Plant.Name},
		},
		Accepter: &api.TradeUser{Surname: trade.Accepter.Surname,
			Name:       trade.Accepter.Name,
			FatherName: trade.Accepter.FatherName,
			Plant:      &api.TradePlant{Name: trade.Accepter.Plant.Name},
		},
	}
	return &api.GetTradeOutV1Response{
		Trade: result,
	}, nil
}
