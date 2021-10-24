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
}

func NewRoutine(frequency float64, granularity string) *MyRoutine {
	return &MyRoutine{
		killChannel: make(chan bool),
		Frequency:   frequency,
		Granularity: granularity,
	}
}

func (this *MyRoutine) setKillChannel(killChannel chan bool) {
	this.killChannel = killChannel
}

type RoutineRepository struct {
	nextId   int
	idMutex  *sync.Mutex
	routines map[int]MyRoutine
}

func NewRoutineRepository() *RoutineRepository {
	this := RoutineRepository{}
	this.nextId = 1
	this.routines = make(map[int]MyRoutine)
	this.idMutex = &sync.Mutex{}
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
	fmt.Println(id, routine.Id)
	this.routines[id] = *routine
	return id
}
