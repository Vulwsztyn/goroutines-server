package main

import (
	"fmt"
	"sync"
)

type MyRoutine struct {
	killChannel chan bool
	Frequency   float64 `json:"frequency"`
	Granularity string  `json:"granularity"`
	Id          int     `json:"id"`
	db *Db
}

func NewRoutine(frequency float64, granularity string, db *Db) *MyRoutine {
	return &MyRoutine{
		killChannel: make(chan bool),
		Frequency:   frequency,
		Granularity: granularity,
	}
}

func (this *MyRoutine) setKillChannel(killChannel chan bool) {
	this.killChannel = killChannel
}

func (this *MyRoutine) run() {
	this.db.insertTs(this.Id)
}


type RoutineRepository struct {
	nextId   int
	idMutex  *sync.Mutex
	routines map[int]MyRoutine
	db *Db
}

func NewRoutineRepository(db *Db) *RoutineRepository {
	this := RoutineRepository{}
	this.nextId = 1
	this.routines = make(map[int]MyRoutine)
	this.idMutex = &sync.Mutex{}
	this.db = db
	return &this
}
func (this *RoutineRepository) getRoutines() map[int]MyRoutine {
	return this.routines
}
func (this *RoutineRepository) getRoutine(id int) MyRoutine {
	return this.routines[id]
}
func (this *RoutineRepository) addRoutine(routine *MyRoutine) int {
	this.idMutex.Lock()
	id := this.nextId
	this.nextId++
	this.idMutex.Unlock()
	routine.Id = id
	routine.db = this.db
	fmt.Println(id, routine.Id)
	this.routines[id] = *routine
	return id
}

func (this *RoutineRepository) removeRoutine(id int) {
	_, ok := this.routines[id]
    if ok {
        delete(this.routines, id)
    }
}
