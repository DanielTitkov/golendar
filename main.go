package main

import (
	"github.com/DanielTitkov/golendar/api"
	"github.com/DanielTitkov/golendar/config"
	"github.com/DanielTitkov/golendar/internal"
	"github.com/DanielTitkov/golendar/logger"
)

func mockEvents(s internal.Storage) {
	events := []internal.Event{
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
	s, err := internal.PrepareStorage(c)
	if err != nil {
		l.Fatal(err)
	}

	mockEvents(s)
	l.Infof("Server started. Listening on port %s, using '%s' storage", c.Port, c.Storage)
	api.HTTPHandleRequests(s, l)
}
