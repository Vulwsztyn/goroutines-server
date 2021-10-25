package main

import (
	"fmt"
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
	fmt.Println(routine.Id, time.Now().Format("2006-01-02 15:04:05"))
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
				fmt.Println(routine.Id, time.Now().Format("2006-01-02 15:04:05"))
				time.Sleep(time.Duration(sleepTime) * time.Second)
			}
		}
	}()
}
func (this *AsyncManager) killRoutine(routine MyRoutine) {
	routine.killChannel <- true
}
