package grpc

import (
	"time"

	pb "github.com/DanielTitkov/golendar/api/grpc/golendarpb"
	"github.com/DanielTitkov/golendar/internal/event"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
)

func pbtsToTime(ts *timestamp.Timestamp) time.Time {
	t, err := ptypes.Timestamp(ts)
	if err != nil {
		t = time.Now() // this is stupid but it should not happen
	}
	return t
}

func timeToPBTS(t time.Time) *timestamp.Timestamp {
	ts := &timestamp.Timestamp{}
	ts, _ = ptypes.TimestampProto(t)
	return ts
}

func pbToEvent(e *pb.Event) event.Event {
	eventUUID, _ := uuid.Parse(e.EventUUID)
	return event.Event{
		UUID:     eventUUID,
		Title:    e.Title,
		Datetime: pbtsToTime(e.Datetime),
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
		Datetime:  timeToPBTS(e.Datetime),
		Duration:  e.Duration,
		Desc:      e.Desc,
		User:      e.User,
		Notify:    e.Notify,
	}
}
