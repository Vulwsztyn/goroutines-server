package main

import (
	"time"
)

type AsyncManagerInterface interface {
	runRoutine(routine MyRoutineInterface)
	killRoutine(routine MyRoutineInterface)
}

type AsyncManager struct {
}

func NewAsyncManager() *AsyncManager {
	return &AsyncManager{}
}
func (this *AsyncManager) runRoutine(routine MyRoutineInterface) {
	routine.init()
	go func() {
		sleepTime := routine.getFrequency() * (map[string]float64{
			"second": 1,
			"minute": 60,
			"hour":   3600,
		}[routine.getGranularity()])

		for {
			select {
			case <-routine.getKillChannel():
				return
			default:
				routine.run()
				time.Sleep(time.Duration(sleepTime) * time.Second)
			}
		}
	}()
}
func (this *AsyncManager) killRoutine(routine MyRoutineInterface) {
	routine.kill()
}
