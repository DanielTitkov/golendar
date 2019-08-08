package api

import (
	"net/http"

	"github.com/DanielTitkov/golendar/internal"
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

// HTTPHandleRequests sets up routes and starts server
func HTTPHandleRequests(s internal.Storage, l *zap.SugaredLogger) {
	mux := http.DefaultServeMux
	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		eventEndpointHandler(w, r, s, l)
	})
	l.Fatal(http.ListenAndServe(":3000", requestLogger(mux, l)))
}
