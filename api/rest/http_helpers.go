package rest

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

func parseUUID(URL url.URL) (uuid.UUID, error) {
	var emptyUUID uuid.UUID
	eventUUIDSlice, ok := URL.Query()["UUID"]
	if !ok {
		return emptyUUID, errors.New("UUID is not provided")
	}
	eventUUID, err := uuid.Parse(eventUUIDSlice[0])
	if err != nil {
		return emptyUUID, err
	}
	return eventUUID, nil
}

func httpRespond(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Write([]byte(message))
}
