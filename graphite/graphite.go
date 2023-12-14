package graphite

import (
	"github.com/jtaczanowski/go-graphite-client"
)

var Client *graphite.Client

func NewClient() {
	graphiteClient := graphite.NewClient("localhost", 2003, "projects", "tcp")
	Client = graphiteClient
}
