package app

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"plants/internal/config"
	gw "plants/internal/pb/plantsapi/github.com/moevm/nosql2h24-plants/server/api/plantsapi"
	"plants/internal/service"
	"plants/internal/storage"

	"golang.org/x/sync/errgroup"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"google.golang.org/grpc"
)

func Run(cfg *config.Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	gRPCServer := grpc.NewServer()
	mongoURL := fmt.Sprintf(
		"%v:%v",
		cfg.Mongo.Domen,
		cfg.Mongo.Port,
	)
	storage, err := storage.New(ctx, mongoURL, cfg.Mongo.DataBase)
	if err != nil {
		return err
	}
	plants_api := service.NewAPI(storage)
	plants_api.Register(gRPCServer)

	mux := runtime.NewServeMux()
	err = gw.RegisterPlantsAPIHandlerServer(ctx, mux, plants_api)
	if err != nil {
		return err
	}
	var group errgroup.Group
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.GRPC.Port))

	if err != nil {
		return err
	}
	group.Go(func() error {
		slog.Info("Started grpc server")
		return gRPCServer.Serve(l)
	})

	group.Go(func() error {
		slog.Info("Started http server")
		return http.ListenAndServe(":8881", mux)
	})
	return group.Wait()
}
