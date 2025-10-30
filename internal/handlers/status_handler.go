package handlers

import (
	"net/http"
)

// StatusHandler returns current status for monitored services.
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"status\": \"not-implemented\"}"))
}
