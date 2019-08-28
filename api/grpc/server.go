package grpc

import (
	"context"

	pb "github.com/DanielTitkov/golendar/api/grpc/golendarpb"
	"github.com/DanielTitkov/golendar/internal/storage"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type eventServer struct {
	storage storage.Storage
}

// Not fully functional yet. User support should be added
func (srv *eventServer) GetEvent(ctx context.Context, req *pb.GetEventRequest) (*pb.GetEventResponse, error) {
	events := []*pb.Event{}
	errors := ""
	for _, stringUUID := range req.GetEventUUID() {
		eventUUID, err := uuid.Parse(stringUUID)
		if err != nil {
			errors = errors + "\n" + err.Error()
			continue
		}
		event, err := srv.storage.GetEvent(eventUUID)
		if err != nil {
			errors = errors + "\n" + err.Error()
			continue
		}
		events = append(events, eventToPb(event))
	}
	if errors != "" {
		return &pb.GetEventResponse{Event: events, Status: "errors occured: " + errors}, nil
	}
	return &pb.GetEventResponse{Event: events, Status: "ok"}, nil
}

func (srv *eventServer) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	e := req.GetEvent()
	newEvent, err := srv.storage.CreateEvent(pbToEvent(e))
	if err != nil {
		return &pb.CreateEventResponse{Status: "error occured: " + err.Error()}, nil
	}
	return &pb.CreateEventResponse{Event: eventToPb(newEvent), Status: "ok"}, nil
}

func (srv *eventServer) UpdateEvent(ctx context.Context, req *pb.UpdateEventRequest) (*pb.UpdateEventResponse, error) {
	e := req.GetEvent()
	eventUUID, err := uuid.Parse(req.GetEventUUID())
	if err != nil {
		return &pb.UpdateEventResponse{Status: "error occured: " + err.Error()}, nil
	}
	newEvent, err := srv.storage.UpdateEvent(eventUUID, pbToEvent(e))
	if err != nil {
		return &pb.UpdateEventResponse{Status: "error occured: " + err.Error()}, nil
	}
	return &pb.UpdateEventResponse{Event: eventToPb(newEvent), Status: "ok"}, nil
}

func (srv *eventServer) DeleteEvent(ctx context.Context, req *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error) {
	eventUUID, err := uuid.Parse(req.GetEventUUID())
	if err != nil {
		return &pb.DeleteEventResponse{Status: "error occured: " + err.Error()}, nil
	}
	err = srv.storage.DeleteEvent(eventUUID)
	if err != nil {
		return &pb.DeleteEventResponse{Status: "error occured: " + err.Error()}, nil
	}
	return &pb.DeleteEventResponse{Status: "ok"}, nil
}

// NewGolendarGRPCServer return new grcp server
func NewGolendarGRPCServer(storage storage.Storage) *grpc.Server {
	grpcServer := grpc.NewServer()
	pb.RegisterEventServiceServer(grpcServer, &eventServer{storage: storage})
	return grpcServer
}
