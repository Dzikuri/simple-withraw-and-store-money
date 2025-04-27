package main

import (
	"os"

	"github.com/dzikuri/simple-withdraw-and-store-money/api"
	"github.com/dzikuri/simple-withdraw-and-store-money/util"
	"github.com/rs/zerolog"
)

func main() {

	// NOTE: Setup Logger
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// NOTE: Call Database Connection
	dbPool, err := util.ConnectDB(logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
	}

	defer dbPool.Close()

	// NOTE: Call API
	api := api.NewAPIServe(dbPool, logger)
	api.Serve()
}
