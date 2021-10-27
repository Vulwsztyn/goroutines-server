package main

import (
	"fmt"
	"sync"
)

type MyRoutineInterface interface {
	init()
	run()
	kill()
	getFrequency() float64
	getGranularity() string
	getKillChannel() chan bool
}


type MyRoutine struct {
	killChannel chan bool
	Frequency   float64 `json:"frequency"`
	Granularity string  `json:"granularity"`
	Id          int     `json:"id"`
	db          DbInterface
}

func NewRoutine(frequency float64, granularity string) *MyRoutine {
	return &MyRoutine{
		Frequency:   frequency,
		Granularity: granularity,
	}
}

func (this *MyRoutine) init() {
	channel := make(chan bool)
	this.killChannel = channel
}

func (this *MyRoutine) run() {
	this.db.InsertTs(this.Id)
}

func (this *MyRoutine) kill() {
	this.killChannel <- true
}

func (this *MyRoutine) getFrequency() float64 {
	return this.Frequency
}

func (this *MyRoutine) getGranularity() string {
	return this.Granularity
}

func (this *MyRoutine) getKillChannel() chan bool {
	return this.killChannel
}

type RoutineRepositoryInterface interface {
	getRoutines() map[int]MyRoutine
	getRoutine(id int) (MyRoutine, error)
	addRoutine(routine *MyRoutine) int
	removeRoutine(id int)
}

type RoutineRepository struct {
	nextId   int
	idMutex  *sync.Mutex
	routines map[int]MyRoutine
	db       DbInterface
}

func NewRoutineRepository(db DbInterface) *RoutineRepository {
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
func (this *RoutineRepository) getRoutine(id int) (MyRoutine, error) {
	if val, ok := this.routines[id]; ok {
		return val, nil
	} else {
		return MyRoutine{}, fmt.Errorf("Routine with id %d not found", id)
	}
	
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
