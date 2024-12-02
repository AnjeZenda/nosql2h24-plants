package storage

import (
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Storage) ExportDB(ctx context.Context) (string, error) {
	data := make(map[string]interface{})
	collections := []string{"users", "plants", "trades", "care_rules"}

	for _, col := range collections {
		collectionData, err := readCollection(ctx, s.Client, "plants_market", col)
		if err != nil {
			return "", err
		}
		data[col] = collectionData
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func readCollection(ctx context.Context, client *mongo.Client, dbName, colName string) ([]bson.M, error) {
	collection := client.Database(dbName).Collection(colName)
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}
