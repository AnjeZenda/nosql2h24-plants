package trades

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"plants/internal/models"
	v1 "plants/internal/pkg/pb/trades/v1"
)

const (
	offerer  = "offerer"
	accepter = "accepter"
)

type Storage interface {
	GetTrades(context.Context, primitive.ObjectID, string) ([]models.Trade, error)
}

type Handler struct {
	v1.UnimplementedTradesServer
	storage Storage
}

func New(storage Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}
