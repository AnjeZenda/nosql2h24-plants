package service

import (
	"context"
	"plants/internal/models"
	"plants/internal/pb/plantsapi/github.com/moevm/nosql2h24-plants/server/api/plantsapi"

	"google.golang.org/grpc"
)

type Storage interface {
	GetPlantsWithCareRules(ctx context.Context) ([]*models.Plant, error)
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
