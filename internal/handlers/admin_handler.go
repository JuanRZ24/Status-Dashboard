package handlers

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/google/uuid"
	"status-dashboard/internal/models"
	"status-dashboard/internal/repository"
)


// AdminHandler provides CRUD endpoints for managing services (add/remove/list).
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"admin\": \"not-implemented\"}"))
}


func CreateServiceHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	var input models.Service

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w,"JSON invalido", http.StatusBadRequest)
		return
	}

	if input.Method == "" {
		input.Method = "GET"
	}
	if input.ExpectedStatus == 0 {
		input.ExpectedStatus = 200
	}
	if input.TimeoutMS == 0 {
		input.TimeoutMS = 5000
	}
	if input.IntervalSec == 0 {
		input.IntervalSec = 120
	}
	if input.Retries == 0 {
		input.Retries = 2
	}
	input.Status = models.StatusUnknown
	input.Enabled = true

	input.ID = uuid.New().String()




}