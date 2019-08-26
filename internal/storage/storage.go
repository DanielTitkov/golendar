package storage

import (
	"errors"

	"github.com/DanielTitkov/golendar/config"
	"github.com/DanielTitkov/golendar/internal/event"
	mapstorage "github.com/DanielTitkov/golendar/internal/storage/map"

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
}

// PrepareStorage setups storage based on config
func PrepareStorage(c config.Config) (Storage, error) {
	switch c.Storage {
	case "MapStorage":
		s := mapstorage.MapStorage{}
		s.Init()
		return &s, nil
	default:
		return nil, errors.New("unknown storage type: " + c.Storage)
	}
}
