package infras

import (
	"commons/otel/db"
	"database/sql"
	_ "github.com/lib/pq"
	"order-ms/internal/infras/config"
)

type Infra struct {
	DB *sql.DB
}

func Load(cfg *config.Config) (*Infra, error) {
	dbCfg := cfg.DBConfig()

	sqlDB, err := db.NewSqlDB(dbCfg.Driver, dbCfg.GetDataSourceName(), dbCfg.DatabaseName, cfg.DBMaxConnections, cfg.DBMaxIdleConnections)
	if err != nil {
		return nil, err
	}

	return &Infra{
		DB: sqlDB,
	}, nil
}
