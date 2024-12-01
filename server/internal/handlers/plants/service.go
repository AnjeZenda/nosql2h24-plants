package plants

import (
	"context"

	"plants/internal/models"
	v1 "plants/internal/pkg/pb/plants/v1"
)

type Storage interface {
	GetPlantsWithCareRules(ctx context.Context) ([]*models.Plant, error)
	CreateNewCareRule(ctx context.Context, rule *models.CareRules) error
	GetCareRulesForPlant(ctx context.Context, species string) (*models.CareRules, error)
	GetPlants(ctx context.Context, fltr *models.Filter) ([]*models.Plant, error)
	AddPlant(ctx context.Context, plant *models.Plant) error
}

type Handler struct {
	v1.UnimplementedPlantsAPIServer
	storage Storage
}

func New(storage Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}
