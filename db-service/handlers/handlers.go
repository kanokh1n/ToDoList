package handlers

import (
	"db-service/models"
	"db-service/repository"
	"db-service/validation"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateTaskHandler(repo *repository.TaskRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var task models.Task

		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON"})
			return
		}

		if err := repo.Create(ctx, &task); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func GetAllTasksHandler(repo *repository.TaskRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tasks, err := repo.GetAllTasks(ctx)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get all tasks"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tasks)
	}
}

func DeleteTaskHandler(repo *repository.TaskRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)
		taskId, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
			return
		}

		if err := validation.ValidateTaskId(taskId); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		if err := repo.DeleteTask(ctx, taskId); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete task"})
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func CompleteTaskHandler(repo *repository.TaskRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)
		taskId, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
			return
		}

		if err := validation.ValidateTaskId(taskId); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		if err := repo.CompleteTask(ctx, taskId); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to complete task"})
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func GetTaskByIdHandler(repo *repository.TaskRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)
		taskId, err := strconv.Atoi(vars["id"])

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
			return
		}

		if err := validation.ValidateTaskId(taskId); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		task, err := repo.GetTaskById(ctx, taskId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get task by id"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(task)
	}
}
