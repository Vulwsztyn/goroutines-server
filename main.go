package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	routineRepository := NewRoutineRepository()
	asyncManager := NewAsyncManager()
	server := NewServer(routineRepository, asyncManager)
	r := mux.NewRouter()
	r.HandleFunc("/create-worker", server.CreateWorker).Methods("POST")
	r.HandleFunc("/kill-worker", server.KillWorker).Methods("POST")
	http.Handle("/", r)
	fmt.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
