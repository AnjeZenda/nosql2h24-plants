package service

import (
	"context"
	api "plants/internal/pb/plantsapi/github.com/moevm/nosql2h24-plants/server/api/plantsapi"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Implementation) GetCareRuleV1(
	ctx context.Context,
	req *api.GetCareRuleV1Request,
) (*api.GetCareRuleV1Response, error) {
	rules, err := s.storage.GetCareRulesForPlant(ctx, req.Species)
	if err != nil {
		return nil, status.Error(codes.Internal, "cant get plant")
	}
	result := make([]*api.CareRule, 0, len(rules.Description))
	for i, r := range rules.Description {
		result[i] = &api.CareRule{
			User:        r.UserID.Hex(),
			Description: r.DescriptionAddition,
			CreatedAt:   timestamppb.New(r.CreatedAt),
		}
	}
	return &api.GetCareRuleV1Response{
		CareRules: result,
	}, nil
}
