package config

import (
	"commons/db/postgre"
	"fmt"
	"os"
	"time"

	env "github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	AppID      string `env:"APP_ID" envDefault:""`
	AppName    string `env:"APP_NAME" envDefault:"product-service"`
	AppVersion string `env:"APP_VERSION" envDefault:"1.0.0"`

	HTTPPort        string        `env:"HTTP_PORT" envDefault:"8080"`
	IdleTimeout     time.Duration `env:"IDLE_TIMEOUT" envDefault:"5s"`
	ReadTimeout     time.Duration `env:"READ_TIMEOUT" envDefault:"5s"`
	WriteTimeout    time.Duration `env:"WRITE_TIMEOUT" envDefault:"5s"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"10s"`

	DBDriver             string `env:"DB_DRIVER" envDefault:"postgres"`
	DBHost               string `env:"DB_HOST" envDefault:"localhost"`
	DBPort               string `env:"DB_PORT" envDefault:"5432"`
	DBUsername           string `env:"DB_USER" envDefault:"postgres"`
	DBPassword           string `env:"DB_PASSWORD" envDefault:"postgres"`
	DBName               string `env:"DB_NAME" envDefault:"product-db"`
	DBMaxConnections     int    `env:"DB_MAX_CONNS" envDefault:"10"`
	DBMaxIdleConnections int    `env:"DB_MAX_IDLE" envDefault:"10"`
}

func (c *Config) Addr() string {
	return fmt.Sprintf(":%s", c.HTTPPort)
}

func (c *Config) DBConfig() *postgre.Config {
	return &postgre.Config{
		Driver:             c.DBDriver,
		Host:               c.DBHost,
		Port:               c.DBPort,
		Username:           c.DBUsername,
		Password:           c.DBPassword,
		DatabaseName:       c.DBName,
		MaxConnections:     c.DBMaxConnections,
		MaxIdleConnections: c.DBMaxIdleConnections,
	}
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
