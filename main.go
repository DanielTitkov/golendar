package main

import (
	"context"
	"net"
	"sync"

	"github.com/DanielTitkov/golendar/api/grpc"
	"github.com/DanielTitkov/golendar/api/rest"
)

func main() {
	// setup context
	ctx := context.Context(context.Background())

	app := App{}
	err := app.Initialize(ctx, "./logger.json", "./config.yaml")
	if err != nil {
		panic(err)
	}

	err = app.MockEvents()
	if err != nil {
		app.l.Fatalf("Data mocking failed: %v", err)
	}

	var wg sync.WaitGroup

	// start REST API server
	wg.Add(1)
	app.l.Infof("REST Server started. Listening on port %s", app.c.Port)
	go rest.HTTPHandleRequests(app.s, app.l)

	// start GRPC server
	wg.Add(1)
	lis, err := net.Listen("tcp", app.c.Host+":"+app.c.GRPCPort)
	if err != nil {
		app.l.Fatal(err)
	}
	app.l.Infof("GRPC Server started. Listening on %s:%s", app.c.Host, app.c.GRPCPort)
	grpcServer := grpc.NewGolendarGRPCServer(app.s)
	go grpcServer.Serve(lis)

	wg.Wait()
}
