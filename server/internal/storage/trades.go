package storage

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"plants/internal/models"
)

func (s *Storage) GetTrades(ctx context.Context, id primitive.ObjectID, role string) ([]models.Trade, error) {
	collection := s.DataBase.Collection("trades")

	var result []models.Trade
	cur, err := collection.Find(ctx, bson.D{{fmt.Sprintf("%s._id", role), id}})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return []models.Trade{}, nil
		}
		return nil, err
	}

	if err = cur.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Storage) CreateBuyTrade(ctx context.Context, trade *models.Trade) error {
	collection := s.DataBase.Collection("trades")
	_, err := collection.InsertOne(ctx, trade)
	if err != nil {
		return err
	}
	return nil
}
