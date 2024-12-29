package app

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"gcommons/db/postgre"
	"gcommons/otel"
	"gcommons/otel/db"
	"gcommons/otel/trace"

	"user-ms/internal/model"
	"user-ms/internal/pkg/config"
	"user-ms/internal/pkg/protos"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type Application struct {
	cfg    *config.Config
	logger zerolog.Logger
	tracer trace.Tracer

	userRepo UserRepository

	protos.UnimplementedUserServiceServer
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

	return &Application{
		cfg:    cfg,
		logger: logger,
		tracer: otel.GetTracer(),

		userRepo: userDao,
	}, cleanups, nil
}

func (app *Application) InitRestServer() *http.Server {
	return &http.Server{
		Addr:         app.cfg.HTTPAddr,
		Handler:      app.Routes(),
		IdleTimeout:  app.cfg.IdleTimeout,
		ReadTimeout:  app.cfg.ReadTimeout,
		WriteTimeout: app.cfg.WriteTimeout,
	}
}

func (app *Application) InitGRPCServer() (net.Listener, *grpc.Server) {
	grpcSrv := grpc.NewServer()
	protos.RegisterUserServiceServer(grpcSrv, app)

	lis, err := net.Listen("tcp", app.cfg.GRPCAddr)
	if err != nil {
		log.Fatalf("failed to init tcp listener: %v", err)
	}

	return lis, grpcSrv
}
