package main

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/DanielTitkov/golendar/api/grpc"
	"github.com/DanielTitkov/golendar/api/rest"
	"github.com/DanielTitkov/golendar/config"
	"github.com/DanielTitkov/golendar/internal/event"
	"github.com/DanielTitkov/golendar/internal/storage"
	"github.com/DanielTitkov/golendar/logger"
	"github.com/jmoiron/sqlx"
)

func mockEvents(s storage.Storage) error {
	events := []event.Event{
		{Title: "Foo", Desc: "FOOBAR", Notify: true, Datetime: time.Now()},
		{Title: "Spam", Desc: "BAZINGA!", Notify: false, Datetime: time.Now()},
		{Title: "Vookah", User: "Mack", Desc: "You gonna like it", Notify: true, Datetime: time.Now()},
	}
	for _, e := range events {
		_, err := s.CreateEvent(e)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// setup context
	ctx := context.Context(context.Background())

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
	s, err := storage.PrepareStorage(ctx, c)
	if err != nil {
		l.Fatal(err)
	}

	// mock data
	if c.Storage == "Postgres" {
		db, err := sqlx.Open("pgx", c.DBURI)
		if err != nil {
			l.Fatal(err)
		}
		initTableQuery := `
			create table if not exists events (
				id serial primary key,
				UUID text not null,
				title text,
				datetime timestamp,
				duration text,
				description text,
				userid text,
				notify boolean);
		`
		if _, err := db.Exec(initTableQuery); err != nil {
			l.Fatal(err)
		}
		db.Close()
	}

	err = mockEvents(s)
	if err != nil {
		l.Fatalf("Data mocking failed: %v", err)
	}

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
