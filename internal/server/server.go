package server

import (
	"log"
	"net/http"

	"thunderbird.zap/thunderbird-dashboard/internal/docker"
)

type Server struct {
	DockerService *docker.DockerService
}

func New() *Server {
	ds, err := docker.NewDockerService()
	if err != nil {
		log.Fatalf("Error Connecting to Docker: %v", err)
	}
	return &Server{
		DockerService: ds,
	}
}

func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Invoked Home Handler")
	data, err := s.DockerService.ListAllContainers()
	if err != nil {
		log.Printf("Error fetching initial container data: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logContainers(data)
	RenderTemplate(w, BaseFile, data)
}

func (s *Server) StatusRefreshHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Invoked Status Handler")
	data, err := s.DockerService.ListAllContainers()
	if err != nil {
		log.Printf("Error refreshing container data: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logContainers(data)
	RenderTemplate(w, ContainerListFile, data)
}

func logContainers(containers []docker.ContainerInfo) {
	for _, container := range containers {
		log.Printf(
			"ID: %s - Name: %s - Status: %s - Running: %v",
			container.ID[:5],
			container.Name,
			container.Status,
			container.Running,
		)
	}
}
