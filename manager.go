package main

import (
	"sync"
)

type MyRoutine struct {
	killChannel chan int
	Frequency   float64 `json:"frequency"`
	Granularity string  `json:"granularity"`
}

func NewRoutine(frequency float64, granularity string) *MyRoutine {
	return &MyRoutine{
		killChannel: make(chan int),
		Frequency:   frequency,
		Granularity: granularity,
	}
}

type Manager struct {
	nextId   int
	idMutex  *sync.Mutex
	routines map[int]MyRoutine
}

func NewManager() *Manager {
	this := Manager{}
	this.nextId = 1
	this.routines = make(map[int]MyRoutine)
	this.idMutex = &sync.Mutex{}
	return &this
}
func (this *Manager) getRoutines() map[int]MyRoutine {
	return this.routines
}
func (this *Manager) getRoutine(id int) map[int]MyRoutine {
	return map[int]MyRoutine{id:this.routines[id]}
}
func (this *Manager) addRoutine(routine MyRoutine) int {
	this.idMutex.Lock()
	id := this.nextId
	this.nextId++
	this.idMutex.Unlock()
	this.routines[id] = routine
	return id
}
