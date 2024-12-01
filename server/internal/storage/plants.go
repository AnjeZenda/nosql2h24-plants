package storage

import (
	"context"

	"plants/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *Storage) GetPlantsWithCareRules(ctx context.Context) ([]*models.Plant, error) {
	collection := s.DataBase.Collection("plants")
	filter := bson.M{"care_rules": bson.M{"$ne": nil}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	plants := make([]*models.Plant, 0)
	for cursor.Next(ctx) {
		var plant models.Plant
		cursor.Decode(&plant)
		plants = append(plants, &plant)
	}
	return plants, nil
}

func (s *Storage) CreateNewCareRule(ctx context.Context, rule *models.CareRules) error {
	collection := s.DataBase.Collection("care_rules")
	result := collection.FindOne(ctx, bson.M{"species": rule.Species})
	if result.Err() != nil {
		_, err := collection.InsertOne(ctx, rule)
		return err
	}
	newAddition := bson.D{
		{Key: "description_addition", Value: rule.Description[0].DescriptionAddition},
		{Key: "created_at", Value: rule.Description[0].CreatedAt},
		{Key: "user_id", Value: rule.Description[0].UserID},
	}
	update := bson.D{
		{Key: "$push", Value: bson.D{{Key: "description", Value: newAddition}}},
		{Key: "$set", Value: bson.D{{Key: "update_at", Value: rule.UpdatedAt}}},
	}
	_, err := collection.UpdateOne(ctx, bson.M{"species": rule.Species}, update)
	return err
}

func (s *Storage) GetCareRulesForPlant(ctx context.Context, species string) (*models.CareRules, error) {
	collection := s.DataBase.Collection("care_rules")
	filter := bson.M{"species": species}
	cursor := collection.FindOne(ctx, filter)

	var result models.CareRules
	err := cursor.Decode(&result)
	if err != nil {
		return nil, err
	}
	for i, d := range result.Description {
		u, err := s.GetUserById(ctx, d.UserID.Hex())
		if err != nil {
			return nil, err
		}
		result.Description[i].UserName = u.Name
		result.Description[i].UserSurname = u.Surname
		result.Description[i].UserFatherName = u.FatherName
	}
	return &result, nil
}

func (s *Storage) GetPlants(ctx context.Context) ([]*models.Plant, error) {
	collection := s.DataBase.Collection("plants")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	plants := make([]*models.Plant, 0)
	for cursor.Next(ctx) {
		var plant models.Plant
		if err := cursor.Decode(&plant); err != nil {
			return nil, err
		}
		plants = append(plants, &plant)
	}
	return plants, nil
}

func (s *Storage) AddPlant(ctx context.Context, plant *models.Plant) error {
	collection := s.DataBase.Collection("plants")
	_, err := collection.InsertOne(ctx, plant)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetPlantsForTrade(ctx context.Context, id string) ([]*models.Plant, error) {
	collection := s.DataBase.Collection("plants")

	filter := bson.D{
		{"user_id", id},
		{"sold_at", ""},
	}
	doc, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	plants := make([]*models.Plant, 0)
	for doc.Next(ctx) {
		var plant models.Plant
		if err := doc.Decode(&plant); err != nil {
			return nil, err
		}
		plants = append(plants, &plant)
	}
	return plants, nil
}
