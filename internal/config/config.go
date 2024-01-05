package config

import (
	"sync"
	"time"

	"github.com/Falokut/admin_casts_service/internal/repository"
	"github.com/Falokut/admin_casts_service/pkg/jaeger"
	"github.com/Falokut/admin_casts_service/pkg/metrics"
	logging "github.com/Falokut/online_cinema_ticket_office.loggerwrapper"
	"github.com/ilyakaznacheev/cleanenv"
)

type KafkaReaderConfig struct {
	Brokers          []string      `yaml:"brokers"`
	GroupID          string        `yaml:"group_id"`
	ReadBatchTimeout time.Duration `yaml:"read_batch_timeout"`
}

type Config struct {
	LogLevel        string `yaml:"log_level" env:"LOG_LEVEL"`
	HealthcheckPort string `yaml:"healthcheck_port" env:"HEALTHCHECK_PORT"`

	Listen struct {
		Host string `yaml:"host" env:"HOST"`
		Port string `yaml:"port" env:"PORT"`
		Mode string `yaml:"server_mode" env:"SERVER_MODE"` // support GRPC, REST, BOTH
	} `yaml:"listen"`

	PrometheusConfig struct {
		Name         string                      `yaml:"service_name" ENV:"PROMETHEUS_SERVICE_NAME"`
		ServerConfig metrics.MetricsServerConfig `yaml:"server_config"`
	} `yaml:"prometheus"`

	DBConfig     repository.DBConfig `yaml:"db_config"`
	JaegerConfig jaeger.Config       `yaml:"jaeger"`

	KafkaMoviesEventsConfig  KafkaReaderConfig `yaml:"movies_events_kafka"`
	KafkaPersonsEventsConfig KafkaReaderConfig `yaml:"persons_events_kafka"`

	MoviesService struct {
		Addr string `yaml:"addr" env:"MOVIES_SERVICE_ADDRESS"`
	} `yaml:"movies_service"`
	MoviesPersonsService struct {
		Addr string `yaml:"addr" env:"MOVIES_PERSONS_SERVICE_ADDRESS"`
	} `yaml:"movies_persons_service"`
}

var instance *Config
var once sync.Once

const configsPath = "configs/"

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		instance = &Config{}

		if err := cleanenv.ReadConfig(configsPath+"config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Fatal(help, " ", err)
		}
	})

	return instance
}
