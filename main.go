package main

import (
	"github.com/DanielTitkov/golendar/api"
	"github.com/DanielTitkov/golendar/internal"
	lg "github.com/DanielTitkov/golendar/logger"
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
	logger, err := lg.CreateLogger("./logger.json")
	if err != nil {
		panic(err)
	}
	s := internal.MapStorage{}
	s.Init()
	mockEvents(&s)
	logger.Info("Server started")
	api.HTTPHandleRequests(&s, logger)
}
