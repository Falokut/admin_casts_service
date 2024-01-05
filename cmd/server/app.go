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

	appCfg := config.GetConfig()
	logLevel, err := logrus.ParseLevel(appCfg.LogLevel)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Logger.SetLevel(logLevel)

	tracer, closer, err := jaegerTracer.InitJaeger(appCfg.JaegerConfig)
	if err != nil {
		logger.Fatal("cannot create tracer", err)
	}
	logger.Info("Jaeger connected")
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	logger.Info("Metrics initializing")
	metric, err := metrics.CreateMetrics(appCfg.PrometheusConfig.Name)
	if err != nil {
		logger.Fatal(err)
	}

	go func() {
		logger.Info("Metrics server running")
		if err := metrics.RunMetricServer(appCfg.PrometheusConfig.ServerConfig); err != nil {
			logger.Fatal(err)
		}
	}()

	logger.Info("Database initializing")
	database, err := repository.NewPostgreDB(appCfg.DBConfig)
	if err != nil {
		logger.Fatalf("Shutting down, connection to the database is not established: %s", err.Error())
	}

	logger.Info("Repository initializing")
	repo := repository.NewCastsRepository(database, logger.Logger)
	defer repo.Shutdown()

	logger.Info("Healthcheck initializing")
	healthcheckManager := healthcheck.NewHealthManager(logger.Logger,
		[]healthcheck.HealthcheckResource{database}, appCfg.HealthcheckPort, nil)
	go func() {
		logger.Info("Healthcheck server running")
		if err := healthcheckManager.RunHealthcheckEndpoint(); err != nil {
			logger.Fatalf("Shutting down, can't run healthcheck endpoint %s", err.Error())
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		logger.Info("Running movie event consumer")
		movieEventsConsumer := events.NewMovieEventsConsumer(getKafkaReaderConfig(appCfg.KafkaMoviesEventsConfig),
			logger.Logger, repo)
		movieEventsConsumer.Run(ctx)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		logger.Info("Running person event consumer")
		personEventsConsumer := events.NewPersonEventsConsumer(getKafkaReaderConfig(appCfg.KafkaMoviesEventsConfig),
			logger.Logger, repo)
		personEventsConsumer.Run(ctx)
		wg.Done()
	}()

	moviesCheck, err := service.NewMoviesChecker(appCfg.MoviesService.Addr)
	if err != nil {
		logger.Fatal(err)
	}
	personsCheck, err := service.NewPersonsChecker(appCfg.MoviesPersonsService.Addr)
	if err != nil {
		logger.Fatal(err)
	}
	checker := service.NewExistanceChecker(moviesCheck, personsCheck, logger.Logger)

	logger.Info("Service initializing")
	service := service.NewCastsService(logger.Logger, repo, checker)

	logger.Info("Server initializing")
	s := server.NewServer(logger.Logger, service)
	s.Run(getListenServerConfig(appCfg), metric, nil, nil)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGTERM)

	<-quit
	s.Shutdown()
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
