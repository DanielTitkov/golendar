package internal

import "github.com/google/uuid"

// Storage has all methods for storing events in some kind of storage
type Storage interface {
	GetEvent(eventUUID uuid.UUID) (Event, error)
	GetEvents() ([]Event, error)
	GetUserEvents(user string) ([]Event, error)
	CreateEvent(e Event) (Event, error)
	UpdateEvent(eventUUID uuid.UUID, e Event) (Event, error)
	DeleteEvent(eventUUID uuid.UUID) error
}
