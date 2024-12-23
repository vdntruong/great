package postgre

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	Username string `env:"DB_USER" envDefault:"postgres"`
	Password string `env:"DB_PASSWORD" envDefault:"postgres"`
	Database string `env:"DB_NAME" envDefault:"postgres"`

	MaxConnections     int `env:"DB_MAX_CONNS" envDefault:"10"`
	MaxIdleConnections int `env:"DB_MAX_IDLE" envDefault:"10"`
}

func NewDB(cfg *Config) (*sql.DB, func(), error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, nil, err
	}

	if cfg.MaxConnections == 0 {
		cfg.MaxConnections = 10
	}
	db.SetMaxOpenConns(cfg.MaxConnections)

	if cfg.MaxIdleConnections == 0 {
		cfg.MaxIdleConnections = 1
	}
	db.SetMaxIdleConns(cfg.MaxIdleConnections)

	cleanup := func() {
		if db != nil {
			_ = db.Close()
		}
	}

	return db, cleanup, nil
}
