package server

import "net/http"

func (s *Server) SetupRoutes() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", s.HomeHandler)
	http.HandleFunc("/status", s.StatusRefreshHandler)
}
