package storage

import (
	"context"
	"time"

	"plants/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type valsType interface {
	string | primitive.ObjectID
}

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

func (s *Storage) GetPlants(ctx context.Context, fltr *models.Filter) ([]*models.Plant, error) {
	filter := bson.D{{"sold_at", time.Time{}}}
	opts := options.Find()
	opts.SetSort(bson.D{{Key: fltr.SortBy, Value: parseSortType(fltr.IsDesc)}})
	opts.SetLimit(fltr.Size)
	opts.SetSkip((fltr.Page - 1) * fltr.Size)
	if len(fltr.Labels) != 0 {
		filter = append(filter, parseLabelsToBSON(fltr.Labels)...)
	}
	collection := s.DataBase.Collection("plants")
	cursor, err := collection.Find(ctx, filter, opts)
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

func (s *Storage) GetPlant(ctx context.Context, id string) (*models.Plant, error) {
	collection := s.DataBase.Collection("plants")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var plant models.Plant
	if err = collection.FindOne(ctx, bson.D{{"_id", objID}}).Decode(&plant); err != nil {
		return nil, err
	}
	return &plant, nil
}

func (s *Storage) SoldPlant(ctx context.Context, id string) error {
	collection := s.DataBase.Collection("plants")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.UpdateOne(ctx, bson.D{{"_id", objID}}, bson.M{"$set": bson.M{"sold_at": time.Now().UTC()}})
	if err != nil {
		return err
	}
	return nil

}

func (s *Storage) GetPlantsByIds(ctx context.Context, ids []string, fltr *models.Filter) ([]*models.Plant, error) {
	var (
		plants []*models.Plant
	)
	collection := s.DataBase.Collection("plants")
	objIDs, err := convertStringToIDs(ids)
	if err != nil {
		return nil, err
	}
	filter := parseLabelsToBSON(fltr.Labels)
	filter = append(filter, createOrFilter("_id", objIDs)...)
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var plant models.Plant
		if err := cur.Decode(&plant); err != nil {
			return nil, err
		}
		plants = append(plants, &plant)
	}
	return plants, nil
}

func (s *Storage) GetPlantsByUserId(ctx context.Context, userId string, isSold bool) ([]*models.Plant, error) {
	var plants []*models.Plant
	objID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	collection := s.DataBase.Collection("plants")
	filter := bson.D{{"user_id", objID}}
	if isSold {
		filter = append(filter, bson.E{"sold_at", bson.M{"$ne": time.Time{}}})
	} else {
		filter = append(filter, bson.E{"sold_at", time.Time{}})
	}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var plant models.Plant
		if err = cur.Decode(&plant); err != nil {
			return nil, err
		}
		plants = append(plants, &plant)
	}
	return plants, nil
}

func parseSortType(isDesc bool) int8 {
	if isDesc {
		return -1
	}
	return 1
}

func parseLabelsToBSON(labels map[string]interface{}) bson.D {
	bsonFltr := bson.D{}
	for k, v := range labels {
		switch vc := v.(type) {
		case string:
			if k == "species" {
				bsonFltr = append(bsonFltr, bson.E{Key: k, Value: bson.M{"$regex": vc, "$options": "i"}})
			} else {
				bsonFltr = append(bsonFltr, bson.E{Key: k, Value: vc})
			}
		case []string:
			bsonFltr = createOrFilter(k, vc)
		}
	}
	if v, ok := labels["price_to"]; ok && v != -1 {
		bsonFltr = append(bsonFltr, bson.E{Key: "price", Value: bson.M{"$lte": v}})
	}
	if v, ok := labels["price_from"]; ok && v != -1 {
		bsonFltr = append(bsonFltr, bson.E{Key: "price", Value: bson.M{"$gte": v}})
	}
	return bsonFltr
}

func createOrFilter[T valsType](k string, vals []T) bson.D {
	bsonFltr := bson.D{}
	listFltr := make([]bson.D, len(vals))
	for i, e := range vals {
		listFltr[i] = bson.D{{Key: k, Value: e}}
	}
	bsonFltr = append(bsonFltr, bson.E{Key: "$or", Value: listFltr})
	return bsonFltr
}

func convertStringToIDs(ids []string) ([]primitive.ObjectID, error) {
	objIDs := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objIDs[i] = objID
	}
	return objIDs, nil
}
