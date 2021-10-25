package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	routineRepository *RoutineRepository
	asyncManager      *AsyncManager
	db                *Db
}

func NewServer(routineRepository *RoutineRepository, asyncManager *AsyncManager, db *Db) *Server {
	return &Server{routineRepository, asyncManager, db}
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
	this.asyncManager.runRoutine(&routine)
	id := this.routineRepository.addRoutine(&routine)
	response := this.routineRepository.getRoutine(id)
	json, err := json.Marshal(map[int]MyRoutine{id: response})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(json))
}

func (this *Server) KillWorker(w http.ResponseWriter, req *http.Request) {
	request := struct {
		Id int
	}{}
	err := json.NewDecoder(req.Body).Decode(&request)
	id := request.Id
	routine := this.routineRepository.getRoutine(id)
	this.asyncManager.killRoutine(routine)
	this.routineRepository.removeRoutine(id)
	json, err := json.Marshal(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(json))
}

func (this *Server) GetEntries(w http.ResponseWriter, req *http.Request) {
	request := struct {
		Id int
	}{}
	err := json.NewDecoder(req.Body).Decode(&request)
	id := request.Id
	entries := this.db.GetEntriesForRunner(id)
	json, err := json.Marshal(entries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(json))
}

func (this *Server) GetRoutines(w http.ResponseWriter, req *http.Request) {
	routines := this.routineRepository.getRoutines()
	json, err := json.Marshal(routines)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(json))
}
