package handlers

import (
	"bytes"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

var DBServiceURL string

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	resp, err := http.Post(DBServiceURL+"/tasks", "application/json", bytes.NewBuffer(body))
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(DBServiceURL + "/tasks")
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}

func GetTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["id"]

	if taskId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	encodedId := url.PathEscape(taskId)

	resp, err := http.Get(DBServiceURL + "/tasks/" + encodedId)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["id"]

	if taskId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	encodedId := url.PathEscape(taskId)

	req, err := http.NewRequest("DELETE", DBServiceURL+"/tasks/"+encodedId, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}

func CompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["id"]

	if taskId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	encodedId := url.PathEscape(taskId)

	req, err := http.NewRequest("PATCH", DBServiceURL+"/tasks/"+encodedId+"/complete", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}
