package internal

import (
	"errors"

	"github.com/google/uuid"
)

//
type MapStorage struct {
	M map[uuid.UUID]Event
}

//
func (mapStorage *MapStorage) Init() {
	mapStorage.M = make(map[uuid.UUID]Event)
}

//
func (mapStorage *MapStorage) GetEvents() ([]Event, error) {
	events := make([]Event, 0, len(mapStorage.M))
	for _, v := range mapStorage.M {
		events = append(events, v)
	}
	return events, nil
}

//
func (mapStorage *MapStorage) GetEvent(eventUUID uuid.UUID) (Event, error) {
	event, ok := mapStorage.M[eventUUID]
	if !ok {
		return event, errors.New("Event not present")
	}
	return event, nil
}

//
func (mapStorage *MapStorage) GetUserEvents(user string) ([]Event, error) {
	return []Event{}, nil
}

//
func (mapStorage *MapStorage) CreateEvent(e Event) (Event, error) {
	e.UUID = uuid.New()
	mapStorage.M[e.UUID] = e
	return e, nil
}

//
func (mapStorage *MapStorage) UpdateEvent(eventUUID uuid.UUID, e Event) (Event, error) {
	e.UUID = eventUUID
	mapStorage.M[eventUUID] = e
	return e, nil
}

//
func (mapStorage *MapStorage) DeleteEvent(eventUUID uuid.UUID) error {
	delete(mapStorage.M, eventUUID)
	return nil
}
