package v1

import (
	"customer-api/internal/concurrency"
	"customer-api/internal/domain"
	"customer-api/internal/ports"
	"encoding/json"
	"net/http"
)

type CustomerHandler struct {
	service ports.CustomerService
	pool    *concurrency.WorkerPool
}

func NewCustomerHandler(service ports.CustomerService, pool *concurrency.WorkerPool) *CustomerHandler {
	return &CustomerHandler{service: service,
		pool: pool}
}

func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer domain.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	job := concurrency.Job{Customer: customer}
	h.pool.AddJob(job)

	result := <-h.pool.Results()

	w.Header().Set("Content-Type", "application/json")
	if result.Success {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": result.Message})
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": result.Message})
	}
}

func (h *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	customer, err := h.service.GetCustomer(id)
	if err != nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
