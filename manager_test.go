package main

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)
const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
    return math.Abs(a - b) <= float64EqualityThreshold
}
func TestAddRoutine(t *testing.T) {
	manager := NewManager()
	manager.addRoutine(*NewRoutine(2, "second"))
	manager.addRoutine(*NewRoutine(0.5, "minute"))
	manager.addRoutine(*NewRoutine(0.314, "hour"))
	routines := manager.getRoutines()
	expectedRoutines := map[int]MyRoutine{
		1:  *NewRoutine(2, "second"),
		2:  *NewRoutine(0.5, "minute"),
		3:  *NewRoutine(0.314, "hour"),
	}
	for k, v := range expectedRoutines {
		v2 := routines[k]
		if v.Granularity != v2.Granularity {
			t.Errorf("For key %d expected granularity %v, got %v", k, v.Granularity, v2.Granularity)
		}
		if !almostEqual(v.Frequency, v2.Frequency) {
			t.Errorf("For key %d expected frequency %v, got %v", k, v.Frequency, v2.Frequency)
		}
	}
}

func TestRoutineJSONability(t *testing.T) {
	manager := NewManager()
	manager.addRoutine(*NewRoutine(2, "second"))
	manager.addRoutine(*NewRoutine(0.5, "minute"))
	manager.addRoutine(*NewRoutine(0.314, "hour"))
	routines := manager.getRoutines()
	expectedJson := `{"1":{"granularity":"second","frequency":3},"2":{"granularity":"minute","frequency":0.5},"3":{"granularity":"hour","frequency":0.314}}`
	resultJson, err := json.Marshal(routines)
	if err != nil {
		t.Errorf("Error marshalling JSON: %v", err)
	}
	require.JSONEq(t, expectedJson, string(resultJson))
}