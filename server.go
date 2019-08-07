package main

import (
	"fmt"
	"log"

	ev "github.com/DanielTitkov/golendar/internal"
)

func operate(s ev.Storage) {
	e1, _ := s.CreateEvent(ev.Event{Title: "Foo", Desc: "FOOBAR"})
	e2, _ := s.CreateEvent(ev.Event{Title: "Spam", Desc: "BAZINGA!"})
	fmt.Println(s.GetEvent(e1.UUID))
	fmt.Println(s.GetEvent(e2.UUID))
	events, err := s.GetEvents()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(events)
	s.UpdateEvent(e1.UUID, ev.Event{Title: "KKK", Desc: "Your time will come!"})
	s.DeleteEvent(e2.UUID)
	events, err = s.GetEvents()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(events)
}

func main() {
	s := ev.MapStorage{}
	s.Init()

	operate(&s)
}
