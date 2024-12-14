package storage

import (
	"context"
	"plants/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *Storage) GetBuyStats(ctx context.Context, filter map[string]interface{}) (*models.BuyStats, int64, error) {
	var stats models.BuyStats
	collection := s.DataBase.Collection("trades")
	pipeline := []bson.D{
		{
			{
				Key: "$match",
				Value: bson.D{
					{"created_at", bson.D{
						{"$gte", filter["from"]},
						{"$lte", filter["to"]},
					}},
					{"type", "buy"},
				},
			},
		},
		{
			{
				Key: "$group",
				Value: bson.D{
					{"_id", bson.D{
						{
							Key: "date",
							Value: bson.D{
								{"$dateToString", bson.D{
									{"format", "%Y-%m-%d"},
									{"date", "$created_at"},
								}},
							},
						},
					}},
					{"count", bson.D{{"$sum", 1}}},
				},
			},
		},
		{
			{
				Key: "$project",
				Value: bson.D{
					{"_id", 0},
					{"date", "$_id.date"},
					{"count", "$count"},
				},
			},
		},
		{
			{
				Key: "$group",
				Value: bson.D{
					{"_id", nil},
					{"count", bson.D{{"$sum", "$count"}}},
					{"info", bson.D{{"$push", bson.D{
						{"date", "$date"},
						{"count", "$count"},
					}}}},
				},
			},
		},
	}
	cur, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, 0, err
	}
	for cur.Next(ctx) {
		if err = cur.Decode(&stats); err != nil {
			return nil, 0, err
		}
	}
	return &stats, 0, nil
}

func (s *Storage) GetTradeStats(ctx context.Context, filter map[string]interface{}) (*models.TradeStats, int64, error) {
	var stats models.TradeStats
	collection := s.DataBase.Collection("trades")
	pipeline := []bson.D{
		{
			{
				Key: "$match",
				Value: bson.D{
					{"created_at", bson.D{
						{"$gte", filter["from"]},
						{"$lte", filter["to"]},
					}},
					{"type", "trade"},
					{"status", bson.M{"$ne": 0}},
				},
			},
		},
		{
			{
				Key: "$group",
				Value: bson.D{
					{"_id", bson.D{
						{
							Key: "date",
							Value: bson.D{
								{"$dateToString", bson.D{
									{"format", "%Y-%m-%d"},
									{"date", "$created_at"},
								}},
							},
						},
						{
							Key:   "status",
							Value: "$status",
						},
					}},
					{"count", bson.D{{"$sum", 1}}},
				},
			},
		},
		{
			{
				Key: "$project",
				Value: bson.D{
					{"_id", 0},
					{"date", "$_id.date"},
					{"status", "$_id.status"},
					{"count", "$count"},
				},
			},
		},
		{
			{
				Key: "$group",
				Value: bson.D{
					{"_id", nil},
					{"count", bson.D{{"$sum", "$count"}}},
					{"info", bson.D{{"$push", bson.D{
						{"date", "$date"},
						{"status", "$status"},
						{"count", "$count"},
					}}}},
				},
			},
		},
	}
	cur, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, 0, err
	}
	for cur.Next(ctx) {
		if err = cur.Decode(&stats); err != nil {
			return nil, 0, err
		}
	}
	return &stats, 0, nil
}

func (s *Storage) GetPlantsStats(ctx context.Context, filter map[string]interface{}) (*models.PlantsStats, int64, error) {
	var stats models.PlantsStats
	collection := s.DataBase.Collection("plants")
	pipeline := []bson.D{
		{
			{
				Key: "$match",
				Value: bson.D{
					{"created_at", bson.D{
						{"$gte", filter["from"]},
						{"$lte", filter["to"]},
					}},
				},
			},
		},
		{
			{
				Key: "$group",
				Value: bson.D{
					{"_id", bson.D{
						{
							Key: "date",
							Value: bson.D{
								{"$dateToString", bson.D{
									{"format", "%Y-%m-%d"},
									{"date", "$created_at"},
								}},
							},
						},
						{
							Key:   "species",
							Value: "$species",
						},
					}},
					{"count", bson.D{{"$sum", 1}}},
				},
			},
		},
		{
			{
				Key: "$project",
				Value: bson.D{
					{"_id", 0},
					{"date", "$_id.date"},
					{"species", "$_id.species"},
					{"count", "$count"},
				},
			},
		},
		{
			{
				Key: "$group",
				Value: bson.D{
					{"_id", nil},
					{"count", bson.D{{"$sum", "$count"}}},
					{"info", bson.D{{"$push", bson.D{
						{"date", "$date"},
						{"species", "$species"},
						{"count", "$count"},
					}}}},
				},
			},
		},
	}
	cur, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, 0, err
	}
	for cur.Next(ctx) {
		if err = cur.Decode(&stats); err != nil {
			return nil, 0, err
		}
	}
	return &stats, 0, nil
}
