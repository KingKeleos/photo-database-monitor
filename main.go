package main

import (
	"log"

	"github.com/KingKeleos/photo-database-monitor/cmd/server"
	"github.com/KingKeleos/photo-database-monitor/graphite"
)

func main() {
	graphite.NewClient()

	api := &server.API{}
	err := api.Serve()
	if err != nil {
		log.Fatalf("Err: %v", err)
	}
}
