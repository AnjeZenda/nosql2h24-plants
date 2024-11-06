package storage

import (
	"context"
	"log/slog"
	"plants/internal/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Storage struct {
	Client   *mongo.Client
	DataBase *mongo.Database
}

func New(ctx context.Context, uri, db string) (*Storage, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		slog.Error("unable to connect")
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		slog.Error("ping doesnt work")
		return nil, err
	}
	database := client.Database(db)
	return &Storage{
		DataBase: database,
		Client:   client,
	}, nil
}

func (s Storage) Disconnect(ctx context.Context) error {
	return s.Client.Disconnect(ctx)
}

func (s *Storage) GetPlantsWithCareRules(ctx context.Context) ([]*models.Plant, error) {
	collection := s.DataBase.Collection("plants")
	filter := bson.M{"care_rules": bson.M{"ne": nil}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	plants := make([]*models.Plant, 0)
	var plant models.Plant
	for cursor.Next(ctx) {
		cursor.Decode(&plant)
		plants = append(plants, &plant)
	}
	return plants, nil
}
