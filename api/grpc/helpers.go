package grpc

import (
	pb "github.com/DanielTitkov/golendar/api/grpc/golendarpb"
	"github.com/DanielTitkov/golendar/internal/event"
	"github.com/google/uuid"
)

func pbToEvent(e *pb.Event) event.Event {
	eventUUID, _ := uuid.Parse(e.EventUUID)
	return event.Event{
		UUID:     eventUUID,
		Title:    e.Title,
		Datetime: e.Datetime,
		Duration: e.Duration,
		Desc:     e.Desc,
		User:     e.User,
		Notify:   e.Notify,
	}
}

func eventToPb(e event.Event) *pb.Event {
	return &pb.Event{
		EventUUID: e.UUID.String(),
		Title:     e.Title,
		Datetime:  e.Datetime,
		Duration:  e.Duration,
		Desc:      e.Desc,
		User:      e.User,
		Notify:    e.Notify,
	}
}
