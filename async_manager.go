package main

import (
	"time"
)

type AsyncManager struct {
}

func NewAsyncManager() *AsyncManager {
	return &AsyncManager{}
}
func (this *AsyncManager) runRoutine(routine *MyRoutine) {
	killChannel := make(chan bool)
	routine.setKillChannel(killChannel)
	go func() {
		sleepTime := routine.Frequency * (map[string]float64{
			"second": 1,
			"minute": 60,
			"hour":   3600,
		}[routine.Granularity])

		for {
			select {
			case <-killChannel:
				return
			default:
				routine.run()
				time.Sleep(time.Duration(sleepTime) * time.Second)
			}
		}
	}()
}
func (this *AsyncManager) killRoutine(routine MyRoutine) {
	routine.killChannel <- true
}
