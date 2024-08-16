package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"userapi/internal/server-client"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"userapi/internal/configs"
	"userapi/internal/models"
)

func main() {
	cfg, err := configs.ParseAndValidate("configs/config.toml")
	if err != nil {
		log.Fatalf("failed to load configs: %v", err)
	}

	var db *gorm.DB

	retryDuration := 2 * time.Minute
	retryInterval := 5 * time.Second
	timeout := time.After(retryDuration)

	for {
		db, err = gorm.Open(postgres.Open(cfg.Database.URL), &gorm.Config{})
		if err == nil {
			log.Println("Successfully connected to the database.")
			break
		}

		log.Printf("failed to connect to the database: %v. Retrying in %v...", err, retryInterval)

		select {
		case <-time.After(retryInterval):
		case <-timeout:
			log.Fatalf("failed to connect to the database after %v: %v", retryDuration, err)
		}
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}

	repo := server_client.NewRepository(db)
	service := server_client.NewService(repo)
	handler := server_client.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/users", handler.CreateUser).Methods("POST")
	router.HandleFunc("/users", handler.GetUsers).Methods("GET")
	router.HandleFunc("/report", handler.GetUsersBy).Methods("GET")

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)

	shutdownComplete := make(chan struct{})

	go func() {
		log.Println("Starting server on :8081")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :8081: %v\n", err)
		}
	}()

	<-stop

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	go func() {
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Server forced to shutdown: %v", err)
		}
		close(shutdownComplete)
	}()

	select {
	case <-shutdownComplete:
		log.Println("Server stopped gracefully.")
	case <-ctx.Done():
		log.Println("Server shutdown timed out.")
	}
}
