package base

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Docker struct {
	cli *client.Client
}

func NewDocker() (*Docker, error) {
	var (
		err    error
		docker Docker
	)
	docker.cli, err = client.NewClient("tcp://127.0.0.1:1990", "1.18.2", nil, map[string]string{"Content-type": "application/x-tar"})
	if err != nil {
		panic(err)
	}
	return &docker, err
}

func (e *Docker) ContainerList() {
	var err error
	containers, err := e.cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}
