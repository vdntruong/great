package app

import (
	"user-ms/internal/pkg/config"
)

type Application struct {
	cfg *config.Config
}

func NewApplication(cfg *config.Config) (*Application, error) {
	return &Application{cfg: cfg}, nil
}
