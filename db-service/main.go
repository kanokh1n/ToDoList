package main

import (
	"context"
	"database/sql"
	"db-service/handlers"
	"db-service/repository"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func main() {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to PostgreSQL")

	// Redis connection
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	defer redisClient.Close()

	log.Println("Connected to Redis")

	repo := repository.NewTaskRepository(db, redisClient)

	router := mux.NewRouter()

	router.HandleFunc("/tasks", handlers.CreateTaskHandler(repo)).Methods("POST")
	router.HandleFunc("/tasks", handlers.GetAllTasksHandler(repo)).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.GetTaskByIdHandler(repo)).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTaskHandler(repo)).Methods("DELETE")
	router.HandleFunc("/tasks/{id}/complete", handlers.CompleteTaskHandler(repo)).Methods("PATCH")

	port := os.Getenv("SERVICE_PORT")

	log.Printf("DB Service running on port %s", port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}
