package main

import (
	"encoding/json"
	"fmt"
)


func main() {
	manager := NewManager()
	manager.addRoutine(*NewRoutine(2, "second"))
	manager.addRoutine(*NewRoutine(0.5, "minute"))
	manager.addRoutine(*NewRoutine(0.314, "hour"))
	routines := manager.getRoutines()
	fmt.Println(routines)
	json, err := json.Marshal(routines)
	fmt.Println(string(json), err)
}