package main

import (
	"log"

	"github.com/KingKeleos/photo-database-monitor/cmd/server"
	"github.com/KingKeleos/photo-database-monitor/graphite"
	"github.com/KingKeleos/photo-database-monitor/postgres"
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
