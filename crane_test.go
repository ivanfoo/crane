package main

import (
	"testing"
)

func TestSchedule(t *testing.T) {
	runtime := &FakeRuntime{}
	agent := &Agent{runtime}
	spec := &Spec{}

	err := agent.Schedule(spec)
	if err != nil {
		t.Error("Could not run workload based on given spec")
	}
}
