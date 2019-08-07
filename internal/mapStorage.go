package internal

import (
	"errors"

	"github.com/google/uuid"
)

// MapStorage is struct for map storage
type MapStorage struct {
	M map[uuid.UUID]Event
}

// Init setups map in MapStorage
func (mapStorage *MapStorage) Init() {
	mapStorage.M = make(map[uuid.UUID]Event)
}

// GetEvents gets all stored events
func (mapStorage *MapStorage) GetEvents() ([]Event, error) {
	events := make([]Event, 0, len(mapStorage.M))
	for _, v := range mapStorage.M {
		events = append(events, v)
	}
	return events, nil
}

// GetEvent gets event by UUID
func (mapStorage *MapStorage) GetEvent(eventUUID uuid.UUID) (Event, error) {
	event, ok := mapStorage.M[eventUUID]
	if !ok {
		return event, errors.New("Event not present")
	}
	return event, nil
}

// GetUserEvents gets all events attributed to specific user
func (mapStorage *MapStorage) GetUserEvents(user string) ([]Event, error) {
	return []Event{}, nil
}

// CreateEvent generates uuid for event object and saves it to storage
func (mapStorage *MapStorage) CreateEvent(e Event) (Event, error) {
	e.UUID = uuid.New()
	mapStorage.M[e.UUID] = e
	return e, nil
}

// UpdateEvent rewrites event with given UUID
func (mapStorage *MapStorage) UpdateEvent(eventUUID uuid.UUID, e Event) (Event, error) {
	e.UUID = eventUUID
	mapStorage.M[eventUUID] = e
	return e, nil
}

// DeleteEvent deletes event with given UUID from storage
func (mapStorage *MapStorage) DeleteEvent(eventUUID uuid.UUID) error {
	delete(mapStorage.M, eventUUID)
	return nil
}
