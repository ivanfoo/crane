package main

import (
	"context"
	"log"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/namespaces"
	"github.com/containerd/containerd/oci"
)

type Runtime interface {
	Run(spec *Spec) error
}

type FakeRuntime struct{}

type ContainerDRuntime struct {
	client *containerd.Client
}

func NewContainerDRuntime() *ContainerDRuntime {
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		log.Fatal(err)
	}

	return &ContainerDRuntime{
		client: client,
	}
}

func (r *ContainerDRuntime) Run(spec *Spec) error {
	ctx := namespaces.WithNamespace(context.Background(), "foo")
	image, err := r.client.Pull(ctx, "docker.io/library/alpine:latest", containerd.WithPullUnpack)

	if err != nil {
		return err
	}

	container, err := r.client.NewContainer(
		ctx,
		"",
		containerd.WithImage(image),
		containerd.WithNewSnapshot("alpine-snapshot", image),
		containerd.WithNewSpec(oci.WithImageConfig(image)),
	)

	if err != nil {
		return err
	}

	task, err := container.NewTask(ctx, cio.NewCreator(cio.WithStdio))
	if err != nil {
		return err
	}

	if err := task.Start(ctx); err != nil {
		return err
	}

	return nil
}

func (r *FakeRuntime) Run(spec *Spec) error {
	return nil
}
