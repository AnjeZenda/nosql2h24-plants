package service

import (
	"context"
	"plants/internal/models"
	api "plants/internal/pb/plantsapi/github.com/moevm/nosql2h24-plants/server/api/plantsapi"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Implementation) CreateNewCareRuleV1(
	ctx context.Context,
	req *api.CreateNewCareRuleV1Request,
) (*api.CreateNewCareRuleV1Response, error) {
	userId, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}
	err = s.storage.CreateNewCareRule(ctx, &models.CareRules{
		Species:   req.Species,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Description: []models.CareDescriptionPart{
			{
				UserID:              userId,
				DescriptionAddition: req.DescriptionAddition,
				CreatedAt:           time.Now().UTC(),
			},
		},
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error occured")
	}
	return &api.CreateNewCareRuleV1Response{}, nil
}
