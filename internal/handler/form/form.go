package form

import (
	"formapp/internal/db/data"
	"formapp/internal/model"
	"html/template"
	"log"
	"net/http"
)

type UserData struct {
	Name       string
	Email      string
	Number     string
	Address    string
	Education  string
	Experience string
	Skills     string
}

func RenderFormPage(w http.ResponseWriter, r *http.Request) {
	filePath := "web/templates/form.html"
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		http.Error(w, "Error loading content", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		return
	}
}

func HandlerFormSubmission(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		userData := model.UserData{
			Name:       r.FormValue("name"),
			Email:      r.FormValue("email"),
			Number:     r.FormValue("phone"),
			Address:    r.FormValue("address"),
			Education:  r.FormValue("education"),
			Experience: r.FormValue("experience"),
			Skills:     r.FormValue("skills"),
		}
		dbSql, _ := data.ConnectDB()
		if err := data.InsertData(dbSql, userData); err != nil {
			log.Fatalf("Error insert data to database: %v", err)
		}
		renderSubmitPage(w, r)
		return
	}
	http.Error(w, "Method GET is not allowed", http.StatusMethodNotAllowed)
}

func renderSubmitPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/submit.html")
	if err != nil {
		log.Fatalf("Error loading submit-template: %v", err)
	}
	if err := tmpl.Execute(w, nil); err != nil {
		return
	}
}
