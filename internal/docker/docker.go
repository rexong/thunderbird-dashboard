package docker

import (
	"context"

	"github.com/moby/moby/client"
)

const CONTAINER_RUNNING = "running"

type ContainerInfo struct {
	ID      string
	Name    string
	Status  string
	Running bool
}

type DockerService struct {
	cli *client.Client
	ctx context.Context
}

type DockerServicer interface {
	ListAllContainers()
}

func NewDockerService() (*DockerService, error) {
	ctx := context.Background()
	cli, err := client.New(client.FromEnv, client.WithAPIVersionFromEnv())
	if err != nil {
		return nil, err
	}
	return &DockerService{
		cli: cli,
		ctx: ctx,
	}, nil
}

func (s *DockerService) ListAllContainers() ([]ContainerInfo, error) {
	results, err := s.cli.ContainerList(s.ctx, client.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}
	var containers []ContainerInfo
	for _, c := range results.Items {
		container := ContainerInfo{
			ID:      c.ID,
			Name:    c.Names[0][1:],
			Status:  c.Status,
			Running: c.State == CONTAINER_RUNNING,
		}
		containers = append(containers, container)
	}
	return containers, nil
}
