package app

import (
	"context"
	"fmt"
	"net"
	"plants/internal/config"
	gw "plants/internal/pb/plantsapi/github.com/moevm/nosql2h24-plants/server/api/plantsapi"
	"plants/internal/service"

	"golang.org/x/sync/errgroup"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"google.golang.org/grpc"
)

func Run(cfg *config.Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	gRPCServer := grpc.NewServer()
	// mongoURL := fmt.Sprintf(
	// 	"%v%v:%v@%v:%v",
	// 	cfg.Mongo.Domen,
	// 	cfg.Mongo.User,
	// 	cfg.Mongo.Password,
	// 	cfg.Mongo.DataBase,
	// 	cfg.Mongo.Port,
	// )
	// storage, err := storage.New(ctx, mongoURL, cfg.Mongo.DataBase)
	// if err != nil {
	// 	return err
	// }
	plants_api := service.NewAPI(struct{}{})
	plants_api.Register(gRPCServer)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(50000000)),
	}
	err := gw.RegisterPlantsAPIHandlerServer(ctx, mux, plants_api)
	if err != nil {
		return err
	}
	var group errgroup.Group
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.GRPC.Port))

	if err != nil {
		return err
	}
	group.Go(func() error {
		return gRPCServer.Serve(l)
	})

	group.Go(func() error {
		return gw.RegisterPlantsAPIHandlerFromEndpoint(ctx, mux, ":8881", opts)
	})
	return group.Wait()
}
