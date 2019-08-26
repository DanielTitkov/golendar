package main

import (
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
	s, err := storage.PrepareStorage(c)
	if err != nil {
		l.Fatal(err)
	}

	mockEvents(s)
	l.Infof("Server started. Listening on port %s, using '%s' storage", c.Port, c.Storage)
	rest.HTTPHandleRequests(s, l)
}
