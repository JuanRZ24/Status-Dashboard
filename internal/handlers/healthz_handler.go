package handlers

import (
	"net/http"
)

// HealthzHandler is a simple liveness probe for the service.
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
