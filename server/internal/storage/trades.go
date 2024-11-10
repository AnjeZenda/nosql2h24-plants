// Copyright (c) 2023-2024, KNS Group LLC ("YADRO").
// All Rights Reserved.
// This software contains the intellectual property of YADRO
// or is licensed to YADRO from third parties. Use of this
// software and the intellectual property contained therein is expressly
// limited to the terms and conditions of the License Agreement under which
// it is provided by YADRO.
//

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
