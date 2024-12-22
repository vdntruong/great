package config

import (
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	AppName    string `env:"APP_NAME" envDefault:"user-ms-api"`
	AppVersion string `env:"APP_VERSION" envDefault:"1.0.0"`

	Addr string `env:"ADDR" envDefault:":8080"`

	IdleTimeout  time.Duration `env:"IDLE_TIMEOUT" envDefault:"5s"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" envDefault:"5s"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" envDefault:"5s"`

	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"10s"`
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
