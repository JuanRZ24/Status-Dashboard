package handlers

import (
	"net/http"
)

// AdminHandler provides CRUD endpoints for managing services (add/remove/list).
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"admin\": \"not-implemented\"}"))
}
