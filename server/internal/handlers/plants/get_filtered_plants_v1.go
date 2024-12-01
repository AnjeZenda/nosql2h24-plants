package plants

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"plants/internal/models"
	api "plants/internal/pkg/pb/plants/v1"
)

func (h *Handler) GetFilteredPlantsV1(
	ctx context.Context,
	req *api.GetFilteredPlantsV1Request,
) (*api.GetFilteredPlantsV1Response, error) {
	filter := parseFilter(req)
	plants, err := h.storage.GetPlants(ctx, filter)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not get plants")
	}
	result := make([]*api.GetFilteredPlantsV1Response_Plant, len(plants))
	for i, p := range plants {
		result[i] = &api.GetFilteredPlantsV1Response_Plant{
			Id:        p.ID.Hex(),
			Image:     p.Image,
			Species:   p.Species,
			Price:     p.Price,
			CreatedAt: timestamppb.New(p.CreatedAt),
			Place:     p.Place,
		}
	}
	return &api.GetFilteredPlantsV1Response{
		Plants: result,
	}, nil
}

func parseFilter(req *api.GetFilteredPlantsV1Request) *models.Filter {
	var labels = make(map[string]interface{})
	if len(req.Filter.Species) != 0 {
		labels["species"] = req.Filter.Species
	}
	if len(req.Filter.Type) != 0 {
		labels["type"] = req.Filter.Type
	}
	if len(req.Filter.Size) != 0 {
		labels["size"] = req.Filter.Size
	}
	if len(req.Filter.LightCondition) != 0 {
		labels["light_condition"] = req.Filter.LightCondition
	}
	if len(req.Filter.WateringFrequency) != 0 {
		labels["watering_frequency"] = req.Filter.WateringFrequency
	}
	if len(req.Filter.TemperatureRegime) != 0 {
		labels["temperature_regime"] = req.Filter.TemperatureRegime
	}
	if len(req.Filter.CareComplexity) != 0 {
		labels["care_complexity"] = req.Filter.CareComplexity
	}
	if len(req.Filter.Description) != 0 {
		labels["description"] = req.Filter.Description
	}
	if len(req.Filter.Place) != 0 {
		labels["place"] = req.Filter.Place
	}
	if req.Filter.PriceFrom != 0 {
		labels["price_from"] = req.Filter.PriceFrom
	}
	if req.Filter.PriceTo != 0 {
		labels["price_to"] = req.Filter.PriceTo
	}
	return &models.Filter{
		Page:   req.Page,
		Size:   req.Size,
		SortBy: req.Sort,
		Labels: labels,
	}
}
