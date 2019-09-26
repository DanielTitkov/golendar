package mapstorage

import (
	"errors"
	"sync"

	"github.com/DanielTitkov/golendar/internal/event"
	"github.com/google/uuid"
)

// MapStorage is struct for map storage
type MapStorage struct {
	M  map[uuid.UUID]event.Event
	mx *sync.Mutex
}

// Init setups map in MapStorage
func (mapStorage *MapStorage) Init() {
	mapStorage.M = make(map[uuid.UUID]event.Event)
	mapStorage.mx = &sync.Mutex{}
}

// GetEvents gets all stored events
func (mapStorage *MapStorage) GetEvents() ([]event.Event, error) {
	events := make([]event.Event, 0, len(mapStorage.M))
	for _, v := range mapStorage.M {
		events = append(events, v)
	}
	return events, nil
}

// GetEvent gets event by UUID
func (mapStorage *MapStorage) GetEvent(eventUUID uuid.UUID) (event.Event, error) {
	event, ok := mapStorage.M[eventUUID]
	if !ok {
		return event, errors.New("Event not present")
	}
	return event, nil
}

// GetUserEvents gets all events attributed to specific user
func (mapStorage *MapStorage) GetUserEvents(user string) ([]event.Event, error) {
	return []event.Event{}, nil
}

// CreateEvent generates uuid for event object and saves it to storage
func (mapStorage *MapStorage) CreateEvent(e event.Event) (event.Event, error) {
	mapStorage.mx.Lock()
	defer mapStorage.mx.Unlock()
	e.UUID = uuid.New()
	mapStorage.M[e.UUID] = e
	return e, nil
}

// UpdateEvent rewrites event with given UUID
func (mapStorage *MapStorage) UpdateEvent(eventUUID uuid.UUID, e event.Event) (event.Event, error) {
	mapStorage.mx.Lock()
	defer mapStorage.mx.Unlock()
	e.UUID = eventUUID
	mapStorage.M[eventUUID] = e
	return e, nil
}

// DeleteEvent deletes event with given UUID from storage
func (mapStorage *MapStorage) DeleteEvent(eventUUID uuid.UUID) error {
	mapStorage.mx.Lock()
	defer mapStorage.mx.Unlock()
	delete(mapStorage.M, eventUUID)
	return nil
}

// GetUpcomingEvents is not supported by map storage
func (mapStorage *MapStorage) GetUpcomingEvents(interval int) ([]event.Event, error) {
	return []event.Event{}, errors.New("Method is not supported by Map Storage")
}
