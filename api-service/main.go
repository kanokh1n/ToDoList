package main

import (
	"api-service/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var dbServiceUrl = os.Getenv("DB_SERVICE_URL")

func main() {

	handlers.DBServiceURL = dbServiceUrl

	router := mux.NewRouter()

	router.HandleFunc("/tasks", handlers.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/tasks", handlers.GetAllTasksHandler).Methods("GET")
	router.HandleFunc("/tasks/{title}", handlers.GetTaskByTitleHandler).Methods("GET")
	router.HandleFunc("/tasks/{title}", handlers.DeleteTaskHandler).Methods("DELETE")
	router.HandleFunc("/tasks/{title}/complete", handlers.CompleteTaskHandler).Methods("PATCH")

	log.Println("API service running on port " + os.Getenv("SERVICE_PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("SERVICE_PORT"), router); err != nil {
		log.Fatal(err)
	}
}
