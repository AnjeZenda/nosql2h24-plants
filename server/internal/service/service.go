package service

import (
	"context"
	"plants/internal/models"
	"plants/internal/pb/plantsapi/github.com/moevm/nosql2h24-plants/server/api/plantsapi"

	"google.golang.org/grpc"
)

type Storage interface {
	GetPlantsWithCareRules(ctx context.Context) ([]*models.Plant, error)
	CreateNewCareRule(ctx context.Context, rule *models.CareRules) error
	GetCareRulesForPlant(ctx context.Context, species string) (*models.CareRules, error)
	GetPlants(ctx context.Context) ([]*models.Plant, error)
	AddPlant(ctx context.Context, plant *models.Plant) error
	SearchUser(ctx context.Context, login string, password string) (string, int32, error)
	AddUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id string) (*models.User, error)
	GetTrade(ctx context.Context, id string, mode int32) (*models.Trade, error)
}

type Implementation struct {
	plantsapi.UnimplementedPlantsAPIServer
	storage Storage
}

func (i Implementation) Register(server *grpc.Server) {
	plantsapi.RegisterPlantsAPIServer(server, &i)
}

func NewAPI(storage Storage) *Implementation {
	return &Implementation{
		storage: storage,
	}
}
