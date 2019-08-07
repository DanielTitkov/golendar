package internal

import "github.com/google/uuid"

//
type Storage interface {
	GetEvent(eventUUID uuid.UUID) (Event, error)
	GetEvents() ([]Event, error)
	GetUserEvents(user string) ([]Event, error)
	CreateEvent(e Event) (Event, error)
	UpdateEvent(eventUUID uuid.UUID, e Event) (Event, error)
	DeleteEvent(eventUUID uuid.UUID) error
}
