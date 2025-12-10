package server

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

const (
	TemplatesDirectory = "templates"
	BaseFile           = "base.html"
	ContainerListFile  = "container_list.html"
)

var templates *template.Template

func MustParseTemplates() {
	pattern := filepath.Join(TemplatesDirectory, "*.html")
	t, err := template.ParseGlob(pattern)
	if err != nil {
		log.Fatalf("Error Parsing Templates: %v", err)
	}
	templates = t
	log.Println("HTML templates Parsed")
}

func RenderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmplName, data)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
