package server_client_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"userapi/internal/mocks"
	"userapi/internal/models"
	"userapi/internal/server-client"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRouter(handler *server_client.Handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", handler.CreateUser).Methods("POST")
	router.HandleFunc("/users", handler.GetUsers).Methods("GET")
	router.HandleFunc("/report", handler.GetUsersBy).Methods("GET")
	return router
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockRepo.On("CreateUser", mock.AnythingOfType("*models.User")).Return(nil)

	service := server_client.NewService(mockRepo)
	handler := server_client.NewHandler(service)

	recorder := httptest.NewRecorder()
	router := setupRouter(handler)

	user := models.User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	body, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetUsers(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockRepo.On("GetUsers").Return([]models.User{
		{
			ID:            "550e8400-e29b-41d4-a716-446655440000",
			FirstName:     "John",
			LastName:      "Doe",
			Age:           30,
			RecordingTime: time.Now(),
		},
	}, nil)

	service := server_client.NewService(mockRepo)
	handler := server_client.NewHandler(service)

	recorder := httptest.NewRecorder()
	router := setupRouter(handler)

	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var users []models.User
	json.Unmarshal(recorder.Body.Bytes(), &users)
	assert.Equal(t, 1, len(users))
	mockRepo.AssertExpectations(t)
}
