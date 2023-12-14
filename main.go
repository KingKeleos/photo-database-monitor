package main

import (
	"fmt"
	"time"

	"github.com/KingKeleos/photo-database-monitor/cmd/server"
	"github.com/KingKeleos/photo-database-monitor/graphite"
)

func main() {
	graphite.NewClient()

	metricHandler := server.Handler{}

	for range time.Tick(30 * time.Minute) {
		fmt.Printf("Fetching metrics\n")
		metricHandler.UpdateProjects()
		metricHandler.UpdatePeople()
		metricHandler.UpdateParticipantsToProjects()
		metricHandler.UpdatePostsToProject()
		metricHandler.UpdateSocials()
		metricHandler.NewProjects()
		metricHandler.ActiveProjects()
		metricHandler.FinishedProjects()
	}
}
