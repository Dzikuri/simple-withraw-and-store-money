package util

import (
	"context"
	"fmt"
	"time"

	"github.com/dzikuri/simple-withdraw-and-store-money/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

// ConnectDB returns a new connection pool for the configured PostgreSQL
// database. If the connection cannot be established, it returns an error.
//
// The function takes no arguments and returns a *pgxpool.Pool and an error.
//
// The function sets the maximum number of connections in the pool to 10, the
// minimum number of connections to 2, and the maximum lifetime of a connection
// to 5 minutes. It also pings the database to ensure the connection is alive.
//
// If any error occurs, it will be logged and returned as an error.
func ConnectDB(logger zerolog.Logger) (*pgxpool.Pool, error) {
	{
		host := config.ConfigEnv.DBHost
		user := config.ConfigEnv.DBUsername
		port := config.ConfigEnv.DBPort
		password := config.ConfigEnv.DBPassword
		dbname := config.ConfigEnv.DBName
		sslmode := "disable"

		dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslmode)

		pool, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse config")
			return nil, fmt.Errorf("failed to parse config: %v", err)
		}

		// NOTE: Set the maximum number of connections in the pool
		pool.MaxConns = 10
		pool.MinConns = 2
		pool.MaxConnLifetime = 5 * time.Minute

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		dbPool, err := pgxpool.NewWithConfig(ctx, pool)
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to create pool")
			return nil, fmt.Errorf("failed to create pool: %v", err)
		}

		// NOTE: Ping the database to ensure connection is alive
		if err := dbPool.Ping(ctx); err != nil {
			logger.Fatal().Err(err).Msg("Failed to ping database")
			return nil, fmt.Errorf("failed to ping database: %v", err)
		}

		logger.Info().Msg("Connected to database")

		return dbPool, nil
	}

}
