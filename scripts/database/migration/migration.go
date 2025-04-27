package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/dzikuri/simple-withdraw-and-store-money/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
)

func main() {

	// NOTE: Setup Logger
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	if len(os.Args) < 2 {
		log.Fatalf("Missing migration direction: up or down")
	}

	direction := strings.ToLower(os.Args[1])

	db, err := util.ConnectDB(logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to DB")
	}
	defer db.Close()

	switch direction {
	case "up":
		if err := migrateUp(db); err != nil {
			logger.Fatal().Err(err).Msg("Failed to migrate up")
		}
	case "down":
		if err := migrateDown(db); err != nil {
			logger.Fatal().Err(err).Msg("Failed to migrate down")
		}
	default:
		logger.Fatal().Err(err).Msgf("Failed to migrate with unknown direction : %s", direction)
	}
}

func migrateUp(db *pgxpool.Pool) error {
	// NOTE: Setup Logger
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	ctx := context.Background()

	files, err := os.ReadDir("scripts/database/migration")
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to read migration directory")
		return fmt.Errorf("failed to read migration directory: %w", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		name := file.Name()
		if !strings.HasSuffix(name, ".up.sql") {
			continue
		}

		path := "scripts/database/migration/" + name
		query, err := os.ReadFile(path)
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to read migration file")
			return fmt.Errorf("failed to read migration file %s: %w", name, err)
		}

		logger.Info().Msgf("Applying migration: %s", name)
		_, err = db.Exec(ctx, string(query))
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to execute migration")
			return fmt.Errorf("failed to execute migration %s: %w", name, err)
		}
	}

	logger.Info().Msg("Migration up completed")
	return nil
}

func migrateDown(db *pgxpool.Pool) error {
	// NOTE: Setup Logger
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	ctx := context.Background()

	files, err := os.ReadDir("scripts/database/migration")
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to read migration directory")
		return fmt.Errorf("failed to read migration directory: %w", err)
	}

	// Reverse loop to rollback properly
	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]
		if file.IsDir() {
			continue
		}
		name := file.Name()
		if !strings.HasSuffix(name, ".down.sql") {
			continue
		}

		path := "scripts/database/migration/" + name
		query, err := os.ReadFile(path)
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to read migration file")
			return fmt.Errorf("failed to read migration file %s: %w", name, err)
		}

		logger.Info().Msgf("Rolling back migration: %s", name)
		_, err = db.Exec(ctx, string(query))
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to execute rollback")
			return fmt.Errorf("failed to execute rollback %s: %w", name, err)
		}
	}

	logger.Info().Msg("Migration down completed")
	return nil
}
