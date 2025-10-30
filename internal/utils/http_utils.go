package utils

import (
	"encoding/json"
	"net/http"
)

// JSON writes a JSON response with the given status.
func JSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
