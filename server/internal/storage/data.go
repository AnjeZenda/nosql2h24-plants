package storage

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Storage) ImportDB(ctx context.Context, dataJson string) error {
	collections, err := s.Client.Database("plants_market").ListCollectionNames(ctx, bson.D{})
	if err != nil {
		return errors.New("ошибка получения списка коллекций: " + err.Error())
	}
	for _, collectionName := range collections {
		if err := s.Client.Database("plants_market").Collection(collectionName).Drop(ctx); err != nil {
			return errors.New("ошибка очистки коллекции " + collectionName + ": " + err.Error())
		}
	}

	log.Println("База данных успешно очищена.")

	var data map[string][]bson.M
	if err := json.Unmarshal([]byte(dataJson), &data); err != nil {
		return errors.New("не удалось распарсить JSON: " + err.Error())
	}

	for collectionName, documents := range data {
		if len(documents) == 0 {
			continue
		}
		newDoc := make([]interface{}, len(documents))
		for i := range documents {
			newDoc[i] = documents[i]
		}
		collection := s.Client.Database("plants_market").Collection(collectionName)
		_, err := collection.InsertMany(ctx, newDoc)
		if err != nil {
			return errors.New("ошибка вставки данных в коллекцию " + collectionName + ": " + err.Error())
		}
	}

	return nil
}

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
