package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/oklog/oklog/pkg/group"
	"gitlab.com/propetrov/project_example/config"
	privateApi "gitlab.com/propetrov/project_example/pkg/api/private"
	"gitlab.com/propetrov/project_example/pkg/plants"
	"gitlab.com/propetrov/project_example/pkg/repository"
	"gitlab.com/propetrov/project_example/pkg/transport/endpoints/private"
	"gitlab.com/propetrov/project_example/pkg/transport/endpoints/public"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalln("Can't parse config")
	}
	fmt.Println(cfg)

	db, err := gorm.Open("postgres", cfg.DB.GetDSN())
	if err != nil {
		log.Fatalln(err)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// init repository
	repo := repository.NewRepository(db, logger.With(zap.String("component", "repository")))

	// init service
	service := plants.NewService(logger.With(zap.String("component", "service")), repo)

	// init controller
	publicController := public.NewController(logger.With(zap.String("component", "public")), service)
	privateController := private.NewController(logger.With(zap.String("component", "private")), service)

	var g group.Group
	server := &http.Server{
		Addr:    cfg.HttpPort,
		Handler: publicController.Endpoints(),
	}
	grpcListener, _ := net.Listen("tcp", cfg.GrpcPort)
	grpcServer := grpc.NewServer()

	// ctrl+c/ctrl+z and other terminate hotkeys, which kill proccess
	trap := make(chan struct{})
	g.Add(
		func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("%s", sig)
			case <-trap:
				return nil
			}
		},
		func(err error) {
			close(trap)
		},
	)

	// http server (public)
	g.Add(
		func() error {
			logger.Info("run http server", zap.String("transport", "http"), zap.String("port", cfg.HttpPort))
			return server.ListenAndServe()
		}, func(err error) {
			server.Close()
		},
	)

	// grpc server (private)
	g.Add(
		func() error {
			logger.Info("run grpc server", zap.String("transport", "grpc"), zap.String("port", cfg.GrpcPort))
			privateApi.RegisterPlantsApiServer(grpcServer, privateController)
			return grpcServer.Serve(grpcListener)
		}, func(err error) {
			grpcListener.Close()
		},
	)

	logger.Error("Application exit", zap.Error(g.Run()))
}
