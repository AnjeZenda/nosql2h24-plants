package storage

import (
	"context"
	"plants/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *Storage) GetBuyStats(ctx context.Context, filter map[string]interface{}) (*models.BuyStats, error) {
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
				"$sort",
				bson.D{{
					"date",
					1,
				}},
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
		return nil, err
	}
	for cur.Next(ctx) {
		if err = cur.Decode(&stats); err != nil {
			return nil, err
		}
	}
	return &stats, nil
}

func (s *Storage) GetTradeStats(ctx context.Context, filter map[string]interface{}) (*models.TradeStats, error) {
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
									{"date", "$updated_at"},
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
				"$sort",
				bson.D{{
					"date",
					1,
				}},
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
		return nil, err
	}
	for cur.Next(ctx) {
		if err = cur.Decode(&stats); err != nil {
			return nil, err
		}
	}
	pipeline = []bson.D{
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
				"$sort",
				bson.D{{
					"date",
					1,
				}},
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
	cur, err = collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	var addStats models.TradeStats
	for cur.Next(ctx) {
		if err = cur.Decode(&addStats); err != nil {
			return nil, err
		}
	}

	for i, s := range stats.Info {
		if s.Status == 1 {
			for _, a := range addStats.Info {
				if s.Date == a.Date {
					stats.Info[i].Count = a.Count
					break
				}
			}
		}
	}
	return &stats, nil
}

func (s *Storage) GetPlantsStats(ctx context.Context, filter map[string]interface{}) (*models.PlantsStats, error) {
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
						{"species", "$species"},
						{"count", "$count"},
					}}}},
				},
			},
		},
	}
	cur, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		if err = cur.Decode(&stats); err != nil {
			return nil, err
		}
	}
	return &stats, nil
}

func (s *Storage) GetAdsStats(ctx context.Context, filter map[string]interface{}) (*models.AdsStats, error) {
	var stats models.AdsStats
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
				"$sort",
				bson.D{{
					"date",
					1,
				}},
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
		return nil, err
	}
	for cur.Next(ctx) {
		if err = cur.Decode(&stats); err != nil {
			return nil, err
		}
	}
	return &stats, nil
}
