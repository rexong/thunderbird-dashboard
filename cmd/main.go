package main

import (
	"log"
	"net/http"

	"thunderbird.zap/thunderbird-dashboard/internal/server"
)

func main() {
	server.MustParseTemplates()

	srv := server.New()
	srv.SetupRoutes()

	port := ":8888"
	log.Printf("Server starting on http://localhost%s", port)
	s := &http.Server{
		Addr: port,
	}
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v", port, err)
	}
}
