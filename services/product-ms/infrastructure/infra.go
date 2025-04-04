package infrastructure

import (
	"commons/otel/db"
	"database/sql"
	"product-ms/internal/config"
)

type Infra struct {
	Cfg *config.Config
	DB  *sql.DB
}

func Load(cfg *config.Config) (*Infra, error) {
	dbCfg := cfg.DBConfig()
	sqlDB, err := db.NewSqlDB(dbCfg.Driver, dbCfg.GetDataSourceName(), dbCfg.DatabaseName, cfg.DBMaxConnections, cfg.DBMaxIdleConnections)
	if err != nil {
		return nil, err
	}

	return &Infra{
		Cfg: cfg,
		DB:  sqlDB,
	}, nil
}
