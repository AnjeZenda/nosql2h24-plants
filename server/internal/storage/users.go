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

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"plants/internal/models"
)

func (s *Storage) GetUserByLoginAndPassword(
	ctx context.Context,
	login, password string,
) (string, models.Role, error) {
	collection := s.DataBase.Collection("users")

	filter := bson.D{
		{"$or", bson.A{
			bson.D{{"email", login}},
			bson.D{{"phone", login}},
		}},
		{"password", password},
	}

	var result models.User
	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", 0, ErrUserNotFound
		}
		return "", 0, err
	}
	return result.ID.Hex(), result.Role, nil
}

func (s *Storage) CreateUser(ctx context.Context, user models.User) error {
	collection := s.DataBase.Collection("users")
	if _, err := collection.InsertOne(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	collection := s.DataBase.Collection("users")

	filter := bson.D{{"email", email}}

	var result models.User
	var err error
	if err = collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.User{}, ErrUserNotFound
		}
		return models.User{}, err
	}
	return result, nil
}

func (s *Storage) GetUserById(ctx context.Context, id string) (models.User, error) {
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
