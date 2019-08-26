package rest

import (
	"net/http"

	"github.com/DanielTitkov/golendar/internal/storage"
	"github.com/DanielTitkov/golendar/logger"
	"go.uber.org/zap"
)

func requestLogger(handler http.Handler, l *zap.SugaredLogger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logger.LogHTTPVisit(l, r)
			handler.ServeHTTP(w, r)
		})
}

func eventEndpointHandler(h eventHandler, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.httpGetEvents(w, r)
	case "POST":
		h.httpCreateEvent(w, r)
	case "PUT":
		h.httpUpdateEvent(w, r)
	case "DELETE":
		h.httpDeleteEvent(w, r)
	default:
		httpRespond(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// HTTPHandleRequests sets up routes and starts server
func HTTPHandleRequests(s storage.Storage, l *zap.SugaredLogger) {
	eh := eventHandler{s, l}
	mux := http.DefaultServeMux
	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) { eventEndpointHandler(eh, w, r) })
	l.Fatal(http.ListenAndServe(":3000", requestLogger(mux, l)))
}
