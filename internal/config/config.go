package config

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"sync"
	"time"

	"github.com/Falokut/admin_casts_service/internal/repository"
	"github.com/Falokut/admin_casts_service/pkg/jaeger"
	"github.com/Falokut/admin_casts_service/pkg/logging"
	"github.com/Falokut/admin_casts_service/pkg/metrics"
	"github.com/ilyakaznacheev/cleanenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type DialMethod = string

const (
	Insecure                 DialMethod = "INSECURE"
	InsecureSkipVerify       DialMethod = "INSECURE_SKIP_VERIFY"
	ClientWithSystemCertPool DialMethod = "CLIENT_WITH_SYSTEM_CERT_POOL"
)

type ConnectionSecureConfig struct {
	Method DialMethod `yaml:"dial_method"`
	// Only for client connection with system pool
	ServerName string `yaml:"server_name"`
}

func (c ConnectionSecureConfig) GetGrpcTransportCredentials() (grpc.DialOption, error) {
	if c.Method == Insecure {
		return grpc.WithTransportCredentials(insecure.NewCredentials()), nil
	}

	if c.Method == InsecureSkipVerify {
		return grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})), nil
	}

	if c.Method == ClientWithSystemCertPool {
		certPool, err := x509.SystemCertPool()
		if err != nil {
			return grpc.EmptyDialOption{}, err
		}
		return grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(certPool, c.ServerName)), nil
	}

	return nil, errors.ErrUnsupported
}

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

	KafkaEventsConfig KafkaReaderConfig `yaml:"kafka_events_config"`

	MoviesService struct {
		Addr             string                 `yaml:"addr" env:"MOVIES_SERVICE_ADDRESS"`
		ConnectionConfig ConnectionSecureConfig `yaml:"connection_config"`
	} `yaml:"movies_service"`
	MoviesPersonsService struct {
		Addr             string                 `yaml:"addr" env:"MOVIES_PERSONS_SERVICE_ADDRESS"`
		ConnectionConfig ConnectionSecureConfig `yaml:"connection_config"`
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
