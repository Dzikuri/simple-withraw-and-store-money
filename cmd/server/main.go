package main

import (
	"log"

	"github.com/dzikuri/simple-withdraw-and-store-money/api"
	"github.com/dzikuri/simple-withdraw-and-store-money/util"
)

func main() {
	// NOTE: Call Database Connection
	dbPool, err := util.ConnectDB()
	// TODO: Need to use Logrus or ZeroLogger
	if err != nil {
		log.Fatal(err)
	}

	defer dbPool.Close()

	// NOTE: Call API
	api := api.NewAPIServe(dbPool)
	api.Serve()
}
