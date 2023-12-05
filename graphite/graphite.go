package graphite

import (
	"github.com/jtaczanowski/go-graphite-client"
)

var Client *graphite.Client

func NewClient() {
	graphiteClient := graphite.NewClient("heimserver", 2003, "projects.projects", "tcp")
	Client = graphiteClient
}
