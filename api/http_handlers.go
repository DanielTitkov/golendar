package api

import (
	"encoding/json"
	"net/http"

	"github.com/DanielTitkov/golendar/internal"
	"github.com/DanielTitkov/golendar/logger"
	"go.uber.org/zap"
)

// HTTPHandleRequests sets up routes and starts server
func HTTPHandleRequests(s internal.Storage, l *zap.SugaredLogger) {
	mux := http.DefaultServeMux
	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		eventEndpointHandler(w, r, s, l)
	})
	l.Fatal(http.ListenAndServe(":3000", requestLogger(mux, l)))
}

func requestLogger(handler http.Handler, l *zap.SugaredLogger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logger.LogHTTPVisit(l, r)
			handler.ServeHTTP(w, r)
		})
}

func eventEndpointHandler(w http.ResponseWriter, r *http.Request, s internal.Storage, l *zap.SugaredLogger) {
	switch r.Method {
	case "GET":
		httpGetEvents(w, r, s, l)
	case "POST":
		httpCreateEvent(w, r, s, l)
	case "PUT":
		httpUpdateEvent(w, r, s, l)
	case "DELETE":
		httpDeleteEvent(w, r, s, l)
	default:
		httpRespond(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func httpCreateEvent(w http.ResponseWriter, r *http.Request, s internal.Storage, l *zap.SugaredLogger) {
	decoder := json.NewDecoder(r.Body)
	var e internal.Event
	if err := decoder.Decode(&e); err != nil {
		l.Errorf("JSON decoding failed: %v", err)
		return
	}
	newEvent, err := s.CreateEvent(e)
	if err != nil {
		l.Errorf("event creation failed: %v", err)
		return
	}

	l.Infof("created event with UUID: %v", newEvent.UUID)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(newEvent)
	if err != nil {
		l.Errorf("JSON encoding failed: %v", err)
		return
	}
}

func httpUpdateEvent(w http.ResponseWriter, r *http.Request, s internal.Storage, l *zap.SugaredLogger) {
	eventUUID, err := parseUUID(*r.URL)
	if err != nil {
		httpRespond(w, http.StatusBadRequest, "UUID is not provided or valid UUID")
		l.Error(err)
		return
	}

	l.Infof("update requested for event %v", eventUUID)

	decoder := json.NewDecoder(r.Body)
	var e internal.Event
	if err := decoder.Decode(&e); err != nil {
		l.Errorf("JSON decoding failed: %v", err)
		return
	}

	updatedEvent, err := s.UpdateEvent(eventUUID, e)
	if err != nil {
		httpRespond(w, http.StatusInternalServerError, "update event failed")
		l.Errorf("update event failed: %v", err)
		return
	}

	l.Infof("updated event with UUID: %v", updatedEvent.UUID)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(updatedEvent)
	if err != nil {
		l.Errorf("JSON encoding failed: %v", err)
		return
	}
}

func httpGetEvents(w http.ResponseWriter, r *http.Request, s internal.Storage, l *zap.SugaredLogger) {
	events, err := s.GetEvents()
	if err != nil {
		l.Error(err)
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(events); err != nil {
		l.Errorf("JSON encoding failed: %v", err)
	}
}

func httpDeleteEvent(w http.ResponseWriter, r *http.Request, s internal.Storage, l *zap.SugaredLogger) {
	eventUUID, err := parseUUID(*r.URL)
	if err != nil {
		httpRespond(w, http.StatusBadRequest, "UUID is not provided or valid UUID")
		l.Error(err)
		return
	}

	l.Infof("delete requested for event %v", eventUUID)

	if err = s.DeleteEvent(eventUUID); err != nil {
		httpRespond(w, http.StatusInternalServerError, "delete event failed")
		l.Errorf("delete event failed: %v", err)
		return
	}
}
