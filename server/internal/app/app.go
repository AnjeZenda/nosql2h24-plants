package app

import (
	"fmt"
	"net"
	"plants/internal/service"

	"google.golang.org/grpc"
)

type App struct {
	API  *grpc.Server
	port string
}

func New(port string) *App {
	gRPCServer := grpc.NewServer()
	plants_api := service.NewAPI(struct{}{})
	plants_api.Register(gRPCServer)

	return &App{
		API:  gRPCServer,
		port: port,
	}
}

func (a *App) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", a.port))
	if err != nil {
		return fmt.Errorf("Dead")
	}

	if err = a.API.Serve(l); err != nil {
		return fmt.Errorf("Killed")
	}
	return nil
}
