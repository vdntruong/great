package app

import (
	"fmt"
	"os"

	"commons/db/postgre"
	"commons/otel"
	"commons/otel/db"
	"commons/otel/trace"
	"commons/protos/userpb"

	"user-ms/internal/model"
	"user-ms/internal/pkg/config"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Application struct {
	cfg    *config.Config
	logger zerolog.Logger
	tracer trace.Tracer

	userRepo UserRepository

	userpb.UnimplementedUserServiceServer
	grpc_health_v1.UnimplementedHealthServer
}

func NewApplication(cfg *config.Config) (*Application, []func(), error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().
		Str("app-name", cfg.AppName).
		Str("app-version", cfg.AppVersion).
		Logger()

	var cleanups []func()

	// infrastructures
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

	// db and handlers/services
	userDao := model.NewDAOUser(userDB, otel.GetTracer())

	app := &Application{
		cfg:    cfg,
		logger: logger,
		tracer: otel.GetTracer(),

		userRepo: userDao,
	}
	return app, cleanups, nil
}
