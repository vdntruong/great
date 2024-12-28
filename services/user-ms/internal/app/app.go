package app

import (
	"context"
	"fmt"
	"os"

	"gcommons/db/postgre"
	"gcommons/otel"
	"gcommons/otel/db"
	"gcommons/otel/trace"

	"user-ms/internal/model"
	"user-ms/internal/pkg/config"

	"github.com/rs/zerolog"
)

type Users interface {
	Insert(ctx context.Context, email, username string, password string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, bool, error)
}

type Application struct {
	cfg    *config.Config
	logger zerolog.Logger
	tracer trace.Tracer

	users Users
}

func NewApplication(cfg *config.Config) (*Application, []func(), error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().
		Str("app-name", cfg.AppName).
		Str("app-version", cfg.AppVersion).
		Logger()

	var cleanups []func()

	dbCfg := postgre.Config{
		Host:               cfg.DBHost,
		Port:               cfg.DBPort,
		Username:           cfg.DBUsername,
		Password:           cfg.DBPassword,
		DatabaseName:       cfg.DBName,
		MaxConnections:     cfg.DBMaxConnections,
		MaxIdleConnections: cfg.DBMaxIdleConnections,
	}
	userDB, cleanup, err := db.NewDB(dbCfg.GetDataSourceName(), dbCfg.DatabaseName, cfg.DBMaxConnections, cfg.DBMaxIdleConnections)
	if err != nil {
		return nil, cleanups, fmt.Errorf("failed to connect to database: %w", err)
	}
	cleanups = append(cleanups, cleanup)

	return &Application{
		cfg:    cfg,
		logger: logger,
		tracer: otel.GetTracer(),

		users: model.NewUserModel(userDB, otel.GetTracer()),
	}, cleanups, nil
}
