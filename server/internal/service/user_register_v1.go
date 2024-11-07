package service

import (
	"context"
	"plants/internal/models"
	api "plants/internal/pb/plantsapi/github.com/moevm/nosql2h24-plants/server/api/plantsapi"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Implementation) UserRegister(
	ctx context.Context,
	req *api.UserRegisterV1Request,
) (*api.UserRegisterV1Response, error) {
	name := strings.Fields(req.name)
	email := ""
	phone := ""
	if strings.Contains(req.Login, "@") {
		email = req.Login
	} else {
		phone = req.Login
	}
	err := s.storage.AddUser(ctx, &models.User{
		ID:          primitive.NewObjectID(),
		Surname:     name[0],
		Name:        name[1],
		FatherName:  name[2],
		PhoneNumber: phone,
		Email:       email,
		Password:    req.Password,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "could not create user")
	}
	return &api.UserRegisterV1Response{}, nil
}
