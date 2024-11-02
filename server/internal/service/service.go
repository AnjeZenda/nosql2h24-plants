package service

import (
	"plants/internal/pb/plantsapi"

	"google.golang.org/grpc"
)

type Storage interface {
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
