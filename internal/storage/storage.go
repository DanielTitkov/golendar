package storage

import (
	"context"
	"errors"

	"github.com/DanielTitkov/golendar/config"
	"github.com/DanielTitkov/golendar/internal/event"
	mapstorage "github.com/DanielTitkov/golendar/internal/storage/map"
	"github.com/DanielTitkov/golendar/internal/storage/pgstorage"

	"github.com/google/uuid"
)

// Storage has all methods for storing events in some kind of storage
type Storage interface {
	GetEvent(eventUUID uuid.UUID) (event.Event, error)
	GetEvents() ([]event.Event, error)
	GetUserEvents(user string) ([]event.Event, error)
	CreateEvent(e event.Event) (event.Event, error)
	UpdateEvent(eventUUID uuid.UUID, e event.Event) (event.Event, error)
	DeleteEvent(eventUUID uuid.UUID) error
	GetUpcomingEvents(interval int) ([]event.Event, error)
}

// PrepareStorage setups storage based on config
func PrepareStorage(ctx context.Context, c config.Config) (Storage, error) {
	switch c.Storage {
	case "MapStorage":
		s := mapstorage.MapStorage{}
		s.Init()
		return &s, nil
	case "Postgres":
		s := pgstorage.PGStorage{URI: c.DBURI, Ctx: ctx}
		err := s.Init()
		return &s, err
	default:
		return nil, errors.New("unknown storage type: " + c.Storage)
	}
}
