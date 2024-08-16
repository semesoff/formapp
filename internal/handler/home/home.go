package home

import (
	"log"
	"net/http"
	"text/template"
)

func HandlerHomePage(w http.ResponseWriter, r *http.Request) {
	renderHomePage(w)
}

func renderHomePage(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("web/templates/home.html")
	if err != nil {
		log.Fatalf("Error loading template of home page: %v", err)
	}
	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatalf("Error writing html-template to ResponseWriter: %v", err)
	}
}
