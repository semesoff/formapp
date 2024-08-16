package main

import (
	"formapp/internal/config"
	"formapp/internal/handler/form"
	"formapp/internal/handler/home"
	"formapp/internal/handler/list"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	cfg, _ := config.LoadConfig()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	r.HandleFunc("/", home.HandlerHomePage).Methods("GET")
	r.HandleFunc("/form", form.RenderFormPage).Methods("GET")
	r.HandleFunc("/submit", form.HandlerFormSubmission).Methods("POST")
	r.HandleFunc("/list", list.HandlerListPage).Methods("GET")
	if err := http.ListenAndServe(cfg.Server.Addr, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
