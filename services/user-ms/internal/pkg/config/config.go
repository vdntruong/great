package config

import (
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	AppID      string `env:"APP_ID" envDefault:""`
	AppName    string `env:"APP_NAME" envDefault:"user-service"`
	AppVersion string `env:"APP_VERSION" envDefault:"1.0.0"`

	HTTPPort        string        `env:"HTTP_PORT" envDefault:"8080"`
	IdleTimeout     time.Duration `env:"IDLE_TIMEOUT" envDefault:"5s"`
	ReadTimeout     time.Duration `env:"READ_TIMEOUT" envDefault:"5s"`
	WriteTimeout    time.Duration `env:"WRITE_TIMEOUT" envDefault:"5s"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"10s"`

	GRPCPort string `env:"GRPC_PORT" envDefault:"9090"`

	// OTLPEndpoint for instrument
	OTLPEndpoint string `env:"OTLP_ENDPOINT" envDefault:"localhost:4317"`

	// Postgre database
	DBHost               string `env:"DB_HOST" envDefault:"localhost"`
	DBPort               string `env:"DB_PORT" envDefault:"5432"`
	DBUsername           string `env:"DB_USER" envDefault:"postgres"`
	DBPassword           string `env:"DB_PASSWORD" envDefault:"postgres"`
	DBName               string `env:"DB_NAME" envDefault:"postgres"`
	DBMaxConnections     int    `env:"DB_MAX_CONNS" envDefault:"10"`
	DBMaxIdleConnections int    `env:"DB_MAX_IDLE" envDefault:"10"`
}

func Load() (*Config, error) {
	var envFile string
	var appEnv = os.Getenv("APP_ENV")
	switch appEnv {
	case "local":
		envFile = ".env.local"
	case "test", "dev", "development":
		envFile = ".env.dev"
	default:
		envFile = ".env"
	}

	if err := godotenv.Load(envFile); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("error loading %s file: %v", envFile, err)
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("error parsing %s file: %v", envFile, err)
	}

	return &cfg, nil
}
