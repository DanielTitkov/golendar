package rest

import (
	"encoding/json"
	"net/http"

	"github.com/DanielTitkov/golendar/internal/event"
	"github.com/DanielTitkov/golendar/internal/storage"
	"go.uber.org/zap"
)

type eventHandler struct {
	storage storage.Storage
	logger  *zap.SugaredLogger
}

func (h *eventHandler) httpCreateEvent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var e event.Event
	if err := decoder.Decode(&e); err != nil {
		h.logger.Errorf("JSON decoding failed: %v", err)
		return
	}
	newEvent, err := h.storage.CreateEvent(e)
	if err != nil {
		h.logger.Errorf("event creation failed: %v", err)
		return
	}

	h.logger.Infof("created event with UUID: %v", newEvent.UUID)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(newEvent)
	if err != nil {
		h.logger.Errorf("JSON encoding failed: %v", err)
		return
	}
}

func (h *eventHandler) httpUpdateEvent(w http.ResponseWriter, r *http.Request) {
	eventUUID, err := parseUUID(*r.URL)
	if err != nil {
		httpRespond(w, http.StatusBadRequest, "UUID is not provided or valid UUID")
		h.logger.Error(err)
		return
	}

	h.logger.Infof("update requested for event %v", eventUUID)

	decoder := json.NewDecoder(r.Body)
	var e event.Event
	if err := decoder.Decode(&e); err != nil {
		h.logger.Errorf("JSON decoding failed: %v", err)
		return
	}

	updatedEvent, err := h.storage.UpdateEvent(eventUUID, e)
	if err != nil {
		httpRespond(w, http.StatusInternalServerError, "update event failed")
		h.logger.Errorf("update event failed: %v", err)
		return
	}

	h.logger.Infof("updated event with UUID: %v", updatedEvent.UUID)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(updatedEvent)
	if err != nil {
		h.logger.Errorf("JSON encoding failed: %v", err)
		return
	}
}

func (h *eventHandler) httpGetEvents(w http.ResponseWriter, r *http.Request) {
	events, err := h.storage.GetEvents()
	if err != nil {
		h.logger.Error(err)
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(events); err != nil {
		h.logger.Errorf("JSON encoding failed: %v", err)
	}
}

func (h *eventHandler) httpDeleteEvent(w http.ResponseWriter, r *http.Request) {
	eventUUID, err := parseUUID(*r.URL)
	if err != nil {
		httpRespond(w, http.StatusBadRequest, "UUID is not provided or valid UUID")
		h.logger.Error(err)
		return
	}

	h.logger.Infof("delete requested for event %v", eventUUID)

	if err = h.storage.DeleteEvent(eventUUID); err != nil {
		httpRespond(w, http.StatusInternalServerError, "delete event failed")
		h.logger.Errorf("delete event failed: %v", err)
		return
	}
}
