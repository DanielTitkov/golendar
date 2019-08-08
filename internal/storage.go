package internal

import (
	"errors"

	"github.com/DanielTitkov/golendar/config"
	"github.com/google/uuid"
)

// Storage has all methods for storing events in some kind of storage
type Storage interface {
	GetEvent(eventUUID uuid.UUID) (Event, error)
	GetEvents() ([]Event, error)
	GetUserEvents(user string) ([]Event, error)
	CreateEvent(e Event) (Event, error)
	UpdateEvent(eventUUID uuid.UUID, e Event) (Event, error)
	DeleteEvent(eventUUID uuid.UUID) error
}

// PrepareStorage setups storage based on config
func PrepareStorage(c config.Config) (Storage, error) {
	switch c.Storage {
	case "MapStorage":
		s := MapStorage{}
		s.Init()
		return &s, nil
	default:
		return nil, errors.New("unknown storage type: " + c.Storage)
	}
}
