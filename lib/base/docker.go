package base

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin/json"
	"io"
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

func (e *Docker) ImageList() []types.ImageSummary {
	imgs, err := e.cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil{ panic(err) }
	return imgs
}

func (e *Docker)ImagePull(imgName string)  {
	events, err := e.cli.ImagePull(context.Background(), imgName, types.ImagePullOptions{})
	if err != nil { panic(err) }
	ret := json.NewDecoder(events)
	type Event struct {
		Status         string `json:"status"`
		Error          string `json:"error"`
		Progress       string `json:"progress"`
		ProgressDetail struct {
			Current int `json:"current"`
			Total   int `json:"total"`
		} `json:"progressDetail"`
	}
	var event *Event
	for {
		if err := ret.Decode(&event); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
	}
}

func (e *Docker) ContainerList() []types.Container {
	containers, err := e.cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil { panic(err) }
	return containers
}

func (e *Docker)  {

}
