package app

import (
	"context"
	"fmt"
	"net"
	"plants/internal/config"
	"plants/internal/service"
	"plants/internal/storage"

	"google.golang.org/grpc"
)

func Run(cfg *config.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)
	defer cancel()
	gRPCServer := grpc.NewServer()
	mongoURL := fmt.Sprintf(
		"%v%v:%v@%v:%v",
		cfg.Mongo.Domen,
		cfg.Mongo.User,
		cfg.Mongo.Password,
		cfg.Mongo.DataBase,
		cfg.Mongo.Port,
	)
	storage, err := storage.New(ctx, mongoURL, cfg.Mongo.DataBase)
	if err != nil {
		return err
	}
	plants_api := service.NewAPI(storage)
	plants_api.Register(gRPCServer)
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.GRPC.Port))
	if err != nil {
		return fmt.Errorf("Dead")
	}

	if err = gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("Killed")
	}
	return nil
}
