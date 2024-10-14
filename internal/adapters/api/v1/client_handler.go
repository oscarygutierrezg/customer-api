package v1

import (
	"customer-api/internal/application"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type ClientHandler struct {
	service *application.ClientService
}

func NewClientHandler(service *application.ClientService) *ClientHandler {
	return &ClientHandler{service: service}
}

func (h *ClientHandler) GetClient(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := vars["id"]

	customer, err := h.service.GetClient(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
