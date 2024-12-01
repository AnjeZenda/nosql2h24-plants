package storage

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"plants/internal/models"
)

func (s *Storage) GetUserByLogin(
	ctx context.Context,
	login string,
) (models.User, error) {
	collection := s.DataBase.Collection("users")

	filter := bson.D{
		{"$or", bson.A{
			bson.D{{"email", login}},
			bson.D{{"phone", login}},
		}}}

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

func (s *Storage) UpdateUser(ctx context.Context,
	id string, name string,
	surname string,
	father_name string,
	email string,
	phone_number string,
	photo string) error {

	collection := s.DataBase.Collection("users")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrUserInvalidId
	}

	filter := bson.M{"_id": objID}
	update := bson.D{
		{"$set", bson.D{
			{"name", name},
			{"surname", surname},
			{"father_name", father_name},
			{"email", email},
			{"phone_number", phone_number},
			{"photo", photo},
		}},
	}
	_, errres := collection.UpdateOne(ctx, filter, update)
	if errres != nil {
		return ErrUserNotUpdated
	}

	return nil
}
