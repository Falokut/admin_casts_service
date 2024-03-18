package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Falokut/admin_casts_service/internal/config"
	"github.com/Falokut/admin_casts_service/internal/events"
	"github.com/Falokut/admin_casts_service/internal/handler"
	"github.com/Falokut/admin_casts_service/internal/moviesservice"
	"github.com/Falokut/admin_casts_service/internal/personsservice"
	"github.com/Falokut/admin_casts_service/internal/repository/postgresrepository"
	"github.com/Falokut/admin_casts_service/internal/service"
	casts_service "github.com/Falokut/admin_casts_service/pkg/admin_casts_service/v1/protos"
	jaegerTracer "github.com/Falokut/admin_casts_service/pkg/jaeger"
	"github.com/Falokut/admin_casts_service/pkg/logging"
	"github.com/Falokut/admin_casts_service/pkg/metrics"
	server "github.com/Falokut/grpc_rest_server"
	"github.com/Falokut/healthcheck"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func initHealthcheck(cfg *config.Config, shutdown chan error, resources []healthcheck.HealthcheckResource) {
	logger := logging.GetLogger()
	logger.Info("Healthcheck initializing")
	healthcheckManager := healthcheck.NewHealthManager(logger.Logger,
		resources, cfg.HealthcheckPort, nil)
	go func() {
		logger.Info("Healthcheck server running")
		if err := healthcheckManager.RunHealthcheckEndpoint(); err != nil {
			logger.Errorf("Shutting down, can't run healthcheck endpoint %v", err)
			shutdown <- err
		}
	}()
}

func initMetrics(cfg *config.Config, shutdown chan error) (metrics.Metrics, error) {
	logger := logging.GetLogger()

	tracer, closer, err := jaegerTracer.InitJaeger(cfg.JaegerConfig)
	if err != nil {
		logger.Errorf("Shutting down, error while creating tracer %v", err)
		return nil, err
	}

	logger.Info("Jaeger connected")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	logger.Info("Metrics initializing")
	metric, err := metrics.CreateMetrics(cfg.PrometheusConfig.Name)
	if err != nil {
		logger.Errorf("Shutting down, error while creating metrics %v", err)
		return nil, err
	}

	go func() {
		logger.Info("Metrics server running")
		if err := metrics.RunMetricServer(cfg.PrometheusConfig.ServerConfig); err != nil {
			logger.Errorf("Shutting down, error while running metrics server %v", err)
			shutdown <- err
		}
	}()

	return metric, nil
}

func main() {
	logging.NewEntry(logging.ConsoleOutput)
	logger := logging.GetLogger()
	cfg := config.GetConfig()

	logLevel, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Logger.SetLevel(logLevel)

	shutdown := make(chan error, 1)
	metric, err := initMetrics(cfg, shutdown)
	if err != nil {
		return
	}

	logger.Info("Database initializing")
	database, err := postgresrepository.NewPostgreDB(&cfg.DBConfig)
	if err != nil {
		logger.Errorf("Shutting down, connection to the database is not established: %s", err.Error())
		return
	}

	logger.Info("Repository initializing")
	repo := postgresrepository.NewCastsRepository(database, logger.Logger)
	defer repo.Shutdown()

	initHealthcheck(cfg, shutdown, []healthcheck.HealthcheckResource{database})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		logger.Info("Running events consumer")
		eventsConsumer := events.NewEventsConsumer(getKafkaReaderConfig(cfg.KafkaEventsConfig),
			logger.Logger, repo)
		eventsConsumer.Run(ctx)
		wg.Done()
	}()

	moviesService, err := moviesservice.NewMoviesService(cfg.MoviesService.Addr,
		cfg.MoviesService.ConnectionConfig, logger.Logger)
	if err != nil {
		logger.Errorf("Shutting down, error while creating movies service %v", err)
		return
	}
	defer moviesService.Shutdown()

	personsService, err := personsservice.NewPersonsService(cfg.MoviesPersonsService.Addr,
		cfg.MoviesPersonsService.ConnectionConfig, logger.Logger)
	if err != nil {
		logger.Errorf("Shutting down, error while creating persons service %v", err)
		return
	}
	defer personsService.Shutdown()

	checker := service.NewExistanceChecker(moviesService, personsService)

	logger.Info("Service initializing")
	s := service.NewCastsService(logger.Logger, repo, checker)
	logger.Info("Handler initializing")
	h := handler.NewCastsService(s)

	logger.Info("Server initializing")
	serv := server.NewServer(logger.Logger, h)
	go func() {
		if serr := serv.Run(getListenServerConfig(cfg), metric, nil, nil); serr != nil {
			logger.Errorf("Shutting down, error while running server endpoint %v", serr)
			shutdown <- serr
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGTERM)

	select {
	case <-quit:
		break
	case <-shutdown:
		break
	}
	serv.Shutdown()
	cancel()
	wg.Wait()
}

func getKafkaReaderConfig(cfg config.KafkaReaderConfig) events.KafkaReaderConfig {
	return events.KafkaReaderConfig{
		Brokers:          cfg.Brokers,
		GroupID:          cfg.GroupID,
		ReadBatchTimeout: cfg.ReadBatchTimeout,
	}
}

func getListenServerConfig(cfg *config.Config) server.Config {
	return server.Config{
		Mode:        cfg.Listen.Mode,
		Host:        cfg.Listen.Host,
		Port:        cfg.Listen.Port,
		ServiceDesc: &casts_service.CastsServiceV1_ServiceDesc,
		RegisterRestHandlerServer: func(ctx context.Context, mux *runtime.ServeMux, service any) error {
			serv, ok := service.(casts_service.CastsServiceV1Server)
			if !ok {
				return errors.New(" can't convert")
			}

			return casts_service.RegisterCastsServiceV1HandlerServer(ctx,
				mux, serv)
		},
	}
}
