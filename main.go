package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	manager := NewManager()
	server := NewServer(manager)
	r := mux.NewRouter()
	r.HandleFunc("/create-worker", server.CreateWorker).Methods("POST")
	http.Handle("/", r)
	fmt.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
