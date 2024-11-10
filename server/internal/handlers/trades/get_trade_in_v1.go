package trades

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	api "plants/internal/pkg/pb/trades/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) GetTradeInV1(
	ctx context.Context,
	req *api.GetTradeInV1Request,
) (*api.GetTradeInV1Response, error) {
	if req.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "id required")
	}
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id: "+err.Error())
	}
	trades, err := h.storage.GetTrades(ctx, id, accepter)
	if err != nil {
		return nil, status.Error(codes.Internal, "cant get trade")
	}

	res := &api.GetTradeInV1Response{Trade: make([]*api.Trade, 0)}
	for _, trade := range trades {
		res.Trade = append(res.Trade, &api.Trade{
			Offerer: &api.Trade_User{
				Name:       trade.Offerer.Name,
				Surname:    trade.Offerer.Surname,
				FatherName: trade.Offerer.FatherName,
				Plant:      &api.Trade_Plant{Name: trade.Offerer.Plant.Name},
			},
			Accepter: &api.Trade_User{
				Name:       trade.Accepter.Name,
				Surname:    trade.Accepter.Surname,
				FatherName: trade.Accepter.FatherName,
				Plant:      &api.Trade_Plant{Name: trade.Accepter.Plant.Name},
			},
			Status: api.Status(trade.Status),
		})
	}
	return res, nil
}
