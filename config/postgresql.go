package config

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

func NewPostgresqlConn() *pgxpool.Pool {
	dsn := "postgres://andhika:andhika123@localhost:5436/mymarket?sslmode=disable"
	pgxConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("Failed to parse postgresl config")
	}

	pgxConfig.MaxConns = 50
	pgxConfig.MinConns = 5
	pgxConfig.MaxConnLifetime = 30 * time.Minute
	pgxConfig.MaxConnIdleTime = 5 * time.Minute
	pgxConfig.HealthCheckPeriod = 1 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		log.Fatal("Failed to create pgx pool")
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal("Failed to ping PostgreSQL database")
	}

	return pool
}
