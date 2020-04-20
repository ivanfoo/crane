package main

type Runtime interface{
	Run(spec *Spec) error
}

type FakeRuntime struct{}

type Agent struct{
	runtime Runtime
}

type Spec struct{}

func (a *Agent) Schedule(spec *Spec) error {
	err := a.runtime.Run(spec)
	return err
}

func (r *FakeRuntime) Run(spec *Spec) error {
	return nil
}
