package storage

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"plants/internal/models"
)

func (s *Storage) GetTrades(ctx context.Context, id primitive.ObjectID, role string) ([]models.Trade, error) {
	collection := s.DataBase.Collection("trades")

	var result []models.Trade
	filter := bson.D{
		{"id", id},
		{"type", "buy"},
		{"$or", bson.A{
			bson.D{{"status", 0}},
			bson.D{{"status", 1}},
		}},
	}
	cur, err := collection.Find(ctx, filter)
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

func (s *Storage) PostBuy(ctx context.Context, trade *models.Trade) error {
	collection := s.DataBase.Collection("trades")
	_, err := collection.InsertOne(ctx, trade)
	if err != nil {
		return err
	}

	var traderes models.Trade
	filter := bson.M{"_id": trade.ID}
	errtr := collection.FindOne(ctx, filter).Decode(&traderes)
	if errtr != nil {
		return ErrTradeNotFound
	}

	collection = s.DataBase.Collection("plants")
	offPlantId := trade.Offerer.Plant.ID

	filter = bson.M{"_id": offPlantId}
	update := bson.D{
		{"$set", bson.D{
			{"sold_at", time.Now().UTC()},
		}},
	}
	_, errres := collection.UpdateOne(ctx, filter, update)
	if errres != nil {
		return ErrTradeNotUpdated
	}
	return nil
}

func (s *Storage) GetUserByIdForTrade(ctx context.Context, id string) (models.User, error) {
	collection := s.DataBase.Collection("users")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, ErrUserInvalidId
	}
	filter := bson.M{"_id": objID}
	var result models.User
	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.User{}, ErrUserNotFound
		}
		return models.User{}, err
	}

	return result, nil
}
func (s *Storage) GetPlantByIdForTrade(ctx context.Context, id string) (models.Plant, error) {
	collection := s.DataBase.Collection("plants")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Plant{}, ErrPlantNotFound
	}
	filter := bson.M{"_id": objID}
	var result models.Plant
	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Plant{}, ErrPlantNotFound
		}
		return models.Plant{}, err
	}

	return result, nil
}

func (s *Storage) UpdateTrade(ctx context.Context,
	id string, status int) error {

	collection := s.DataBase.Collection("trades")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrUserInvalidId
	}

	filter := bson.M{"_id": objID}
	update := bson.D{
		{"$set", bson.D{
			{"status", status},
		}},
	}
	_, errres := collection.UpdateOne(ctx, filter, update)
	if errres != nil {
		return ErrTradeNotUpdated
	}
	var trade models.Trade
	errtr := collection.FindOne(ctx, filter).Decode(&trade)
	if errtr != nil {
		return ErrTradeNotFound
	}

	collection = s.DataBase.Collection("plants")
	offPlantId := trade.Offerer.Plant.ID
	filter = bson.M{"_id": offPlantId}
	update = bson.D{
		{"$set", bson.D{
			{"sold_at", time.Now().UTC()},
		}},
	}
	_, errres = collection.UpdateOne(ctx, filter, update)
	if errres != nil {
		return ErrTradeNotUpdated
	}

	accPlantId := trade.Accepter.Plant.ID
	filter = bson.M{"_id": accPlantId}
	update = bson.D{
		{"$set", bson.D{
			{"sold_at", time.Now().UTC()},
		}},
	}
	_, errres = collection.UpdateOne(ctx, filter, update)
	if errres != nil {
		return ErrTradeNotUpdated
	}

	return nil
}

