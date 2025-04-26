package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/dzikuri/simple-withdraw-and-store-money/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/gommon/log"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Missing migration direction: up or down")
	}

	direction := strings.ToLower(os.Args[1])

	db, err := util.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	switch direction {
	case "up":
		if err := migrateUp(db); err != nil {
			log.Fatalf("Failed to migrate up: %v", err)
		}
	case "down":
		if err := migrateDown(db); err != nil {
			log.Fatalf("Failed to migrate down: %v", err)
		}
	default:
		log.Fatalf("Unknown migration direction: %s", direction)
	}
}

func migrateUp(db *pgxpool.Pool) error {
	ctx := context.Background()

	files, err := os.ReadDir("scripts/database/migration")
	if err != nil {
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
			return fmt.Errorf("failed to read migration file %s: %w", name, err)
		}

		log.Infof("Applying migration: %s", name)
		_, err = db.Exec(ctx, string(query))
		if err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", name, err)
		}
	}

	log.Info("Migration up completed")
	return nil
}

func migrateDown(db *pgxpool.Pool) error {
	ctx := context.Background()

	files, err := os.ReadDir("scripts/database/migration")
	if err != nil {
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
			return fmt.Errorf("failed to read migration file %s: %w", name, err)
		}

		log.Infof("Rolling back migration: %s", name)
		_, err = db.Exec(ctx, string(query))
		if err != nil {
			return fmt.Errorf("failed to execute rollback %s: %w", name, err)
		}
	}

	log.Info("Migration down completed")
	return nil
}
