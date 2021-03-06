package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"

	"github.com/gorilla/mux"
)

func main() {
	db := NewDb(func(driverName, dataSourceName string) (SqlInterface, error) {
		return sql.Open(driverName, dataSourceName)
	})
	routineRepository := NewRoutineRepository(db)
	asyncManager := NewAsyncManager()
	server := NewServer(routineRepository, asyncManager, db)
	r := mux.NewRouter()
	r.HandleFunc("/create-worker", server.CreateWorker).Methods("POST")
	r.HandleFunc("/kill-worker", server.KillWorker).Methods("POST")
	r.HandleFunc("/get-workers", server.GetRoutines).Methods("POST")
	r.HandleFunc("/get-entries", server.GetEntries).Methods("POST")
	http.Handle("/", r)
	fmt.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
