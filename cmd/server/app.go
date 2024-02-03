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
	"github.com/Falokut/admin_casts_service/internal/repository"
	"github.com/Falokut/admin_casts_service/internal/service"
	casts_service "github.com/Falokut/admin_casts_service/pkg/admin_casts_service/v1/protos"
	jaegerTracer "github.com/Falokut/admin_casts_service/pkg/jaeger"
	"github.com/Falokut/admin_casts_service/pkg/metrics"
	server "github.com/Falokut/grpc_rest_server"
	"github.com/Falokut/healthcheck"
	logging "github.com/Falokut/online_cinema_ticket_office.loggerwrapper"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func main() {
	logging.NewEntry(logging.ConsoleOutput)
	logger := logging.GetLogger()
	cfg := config.GetConfig()

	logLevel, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Logger.SetLevel(logLevel)

	tracer, closer, err := jaegerTracer.InitJaeger(cfg.JaegerConfig)
	if err != nil {
		logger.Errorf("Shutting down, error while creating tracer %v", err)
		return
	}
	logger.Info("Jaeger connected")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	logger.Info("Metrics initializing")
	metric, err := metrics.CreateMetrics(cfg.PrometheusConfig.Name)
	if err != nil {
		logger.Errorf("Shutting down, error while creating metrics %v", err)
		return
	}

	shutdown := make(chan error, 1)
	go func() {
		logger.Info("Metrics server running")
		if err := metrics.RunMetricServer(cfg.PrometheusConfig.ServerConfig); err != nil {
			logger.Errorf("Shutting down, error while running metrics server %v", err)
			shutdown <- err
		}
	}()

	logger.Info("Database initializing")
	database, err := repository.NewPostgreDB(cfg.DBConfig)
	if err != nil {
		logger.Errorf("Shutting down, connection to the database is not established: %s", err.Error())
		return
	}

	logger.Info("Repository initializing")
	repo := repository.NewCastsRepository(database, logger.Logger)
	defer repo.Shutdown()

	logger.Info("Healthcheck initializing")
	healthcheckManager := healthcheck.NewHealthManager(logger.Logger,
		[]healthcheck.HealthcheckResource{database}, cfg.HealthcheckPort, nil)
	go func() {
		logger.Info("Healthcheck server running")
		if err := healthcheckManager.RunHealthcheckEndpoint(); err != nil {
			logger.Errorf("Shutting down, can't run healthcheck endpoint %s", err.Error())
			shutdown <- err
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		logger.Info("Running movie event consumer")
		movieEventsConsumer := events.NewMovieEventsConsumer(getKafkaReaderConfig(cfg.KafkaMoviesEventsConfig),
			logger.Logger, repo)
		movieEventsConsumer.Run(ctx)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		logger.Info("Running person event consumer")
		personEventsConsumer := events.NewPersonEventsConsumer(getKafkaReaderConfig(cfg.KafkaMoviesEventsConfig),
			logger.Logger, repo)
		personEventsConsumer.Run(ctx)
		wg.Done()
	}()

	moviesCheck, err := service.NewMoviesChecker(cfg.MoviesService.Addr,
		cfg.MoviesService.ConnectionConfig)
	if err != nil {
		logger.Errorf("Shutting down, error while creating movies checker %v", err)
		return
	}
	defer moviesCheck.Shutdown()

	personsCheck, err := service.NewPersonsChecker(cfg.MoviesPersonsService.Addr,
		cfg.MoviesPersonsService.ConnectionConfig)
	if err != nil {
		logger.Errorf("Shutting down, error while creating persons checker %v", err)
		return
	}
	defer personsCheck.Shutdown()

	checker := service.NewExistanceChecker(moviesCheck, personsCheck, logger.Logger)

	logger.Info("Service initializing")
	service := service.NewCastsService(logger.Logger, repo, checker)

	logger.Info("Server initializing")
	s := server.NewServer(logger.Logger, service)
	go func() {
		if err := s.Run(getListenServerConfig(cfg), metric, nil, nil); err != nil {
			logger.Errorf("Shutting down, error while running server endpoint %s", err.Error())
			shutdown <- err
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
	s.Shutdown()
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

			return casts_service.RegisterCastsServiceV1HandlerServer(context.Background(),
				mux, serv)
		},
	}
}
