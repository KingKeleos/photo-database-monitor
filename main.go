package main

import (
	"log"

	"github.com/kingkeleos/database-monitor/cmd/server"
	"github.com/kingkeleos/database-monitor/graphite"
	"github.com/kingkeleos/database-monitor/postgres"
)

func main() {
	graphite.NewClient()
	postgres.ConnectToPostgres()

	api := &server.API{}
	err := api.Serve()
	if err != nil {
		log.Fatalf("Err: %v", err)
	}
}
