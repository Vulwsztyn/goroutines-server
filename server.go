package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	manager *Manager
}

func NewServer(manager *Manager) *Server {
	return &Server{manager}
}

func (this *Server) CreateWorker(w http.ResponseWriter, req *http.Request) {
	var routine MyRoutine
	err := json.NewDecoder(req.Body).Decode(&routine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if routine.Frequency < 1 {
		http.Error(w, "Frequency must be greater or equal than 1", http.StatusBadRequest)
		return
	}
	if map[string]int{"second": 1, "minute": 1, "hour": 1}[routine.Granularity] != 1 {
		http.Error(w, "Granularity must be one of second, minute, hour", http.StatusBadRequest)
		return
	}
	id := this.manager.addRoutine(routine)
	response := this.manager.getRoutine(id)
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(json))
}
