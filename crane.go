package main

import (
	"fmt"
)

type Agent struct {
	runtime Runtime
}

type Spec struct{}

func (a *Agent) Schedule(spec *Spec) error {
	err := a.runtime.Run(spec)
	return err
}

func main() {
	runtime := NewContainerDRuntime()
	err := runtime.Run(nil)
	if err != nil {
		fmt.Println(err)
	}
}
