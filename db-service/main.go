package main

import (
	"database/sql"
	"db-service/handlers"
	"db-service/repository"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
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

	repo := repository.NewTaskRepository(db)

	router := mux.NewRouter()

	router.HandleFunc("/tasks", handlers.CreateTaskHandler(repo)).Methods("POST")
	router.HandleFunc("/tasks", handlers.GetAllTasksHandler(repo)).Methods("GET")
	router.HandleFunc("/tasks/{title}", handlers.GetTaskByTitleHandler(repo)).Methods("GET")
	router.HandleFunc("/tasks/{title}", handlers.DeleteTaskHandler(repo)).Methods("DELETE")
	router.HandleFunc("/tasks/{title}/complete", handlers.CompleteTaskHandler(repo)).Methods("PATCH")

	port := os.Getenv("SERVICE_PORT")

	log.Printf("DB Service running on port %s", port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}
