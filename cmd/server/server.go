package server

import (
	"fmt"
	"log"
	"net"

	DatabaseMonitor "github.com/kingkeleos/database-monitor/grpc"
	"google.golang.org/grpc"
)

type API struct {
	srv grpc.Server
}

func (a *API) Serve() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8081))
	if err != nil {
		log.Fatalf("Could not start grpc-Server: %v", err)
	}
	a.srv = *grpc.NewServer()
	DatabaseMonitor.RegisterDatabaseMonitorServer(&a.srv, &Handler{})
	fmt.Printf("Started server on: %v\n", lis.Addr())
	return a.srv.Serve(lis)
}
