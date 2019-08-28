package main

import (
	"net"
	"sync"

	"github.com/DanielTitkov/golendar/api/grpc"
	"github.com/DanielTitkov/golendar/api/rest"
	"github.com/DanielTitkov/golendar/config"
	"github.com/DanielTitkov/golendar/internal/event"
	"github.com/DanielTitkov/golendar/internal/storage"
	"github.com/DanielTitkov/golendar/logger"
)

func mockEvents(s storage.Storage) {
	events := []event.Event{
		{Title: "Foo", Desc: "FOOBAR"},
		{Title: "Spam", Desc: "BAZINGA!"},
		{Title: "Vookah", User: "Mack", Desc: "You gonna like it"},
	}
	for _, e := range events {
		s.CreateEvent(e)
	}
}

func main() {
	// setup logger
	l, err := logger.CreateLogger("./logger.json")
	if err != nil {
		panic(err)
	}

	// load config from file
	c, err := config.LoadYamlConfig("./config.yaml")
	if err != nil {
		l.Fatalf("config is not loaded: %v", err)
	}

	// setup storage
	l.Infof("Using '%s' storage", c.Storage)
	s, err := storage.PrepareStorage(c)
	if err != nil {
		l.Fatal(err)
	}

	mockEvents(s)

	var wg sync.WaitGroup

	// start REST API server
	wg.Add(1)
	l.Infof("REST Server started. Listening on port %s", c.Port)
	go rest.HTTPHandleRequests(s, l)

	// start GRPC server
	wg.Add(1)
	lis, err := net.Listen("tcp", c.Host+":"+c.GRPCPort)
	if err != nil {
		l.Fatal(err)
	}
	l.Infof("GRPC Server started. Listening on %s:%s", c.Host, c.GRPCPort)
	grpcServer := grpc.NewGolendarGRPCServer(s)
	go grpcServer.Serve(lis)

	wg.Wait()
}
