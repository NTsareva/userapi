package server_client

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"userapi/internal/models"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.RecordingDate = time.Now().Unix()

	if err := h.service.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) GetUsersBy(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	var startDate, endDate time.Time
	var minAge, maxAge int
	var err error

	if startDateStr := queryParams.Get("start_date"); startDateStr != "" {
		startDate, err = time.Parse(time.RFC3339, startDateStr)
		if err != nil {
			http.Error(w, "Invalid start_date format", http.StatusBadRequest)
			return
		}
	}

	if endDateStr := queryParams.Get("end_date"); endDateStr != "" {
		endDate, err = time.Parse(time.RFC3339, endDateStr)
		if err != nil {
			http.Error(w, "Invalid end_date format", http.StatusBadRequest)
			return
		}
	}

	if minAgeStr := queryParams.Get("min_age"); minAgeStr != "" {
		minAge, err = strconv.Atoi(minAgeStr)
		if err != nil {
			http.Error(w, "Invalid min_age", http.StatusBadRequest)
			return
		}
	}

	if maxAgeStr := queryParams.Get("max_age"); maxAgeStr != "" {
		maxAge, err = strconv.Atoi(maxAgeStr)
		if err != nil {
			http.Error(w, "Invalid max_age", http.StatusBadRequest)
			return
		}
	}

	users, count, err := h.service.GenerateReport(startDate, endDate, minAge, maxAge)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"users": users,
		"count": count,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
