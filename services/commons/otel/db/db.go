package db

//https://github.com/open-telemetry/opentelemetry-go-contrib/issues/5#issuecomment-2062564807

import (
	"database/sql"

	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func NewDB(dsn string, dbname string, maxConn, maxIdleConn int) (*sql.DB, func(), error) {
	db, err := otelsql.Open("postgres", dsn,
		otelsql.WithAttributes(semconv.DBSystemPostgreSQL),
		otelsql.WithDBName(dbname))
	if err != nil {
		return nil, nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, nil, err
	}

	if maxConn == 0 {
		maxConn = 10
	}
	db.SetMaxOpenConns(maxConn)

	if maxIdleConn == 0 {
		maxIdleConn = 1
	}
	db.SetMaxIdleConns(maxIdleConn)

	cleanup := func() {
		if db != nil {
			_ = db.Close()
		}
	}

	return db, cleanup, nil
}

func NewSqlDB(driver string, dsn string, dbname string, maxConn, maxIdleConn int) (*sql.DB, error) {
	db, err := otelsql.Open(
		driver, dsn,
		otelsql.WithAttributes(semconv.DBSystemPostgreSQL),
		otelsql.WithDBName(dbname),
	)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)

	return db, nil
}
