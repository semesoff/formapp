package list

import (
	"formapp/internal/db/data"
	"formapp/internal/model"
	"html/template"
	"log"
	"net/http"
)

func getListUsers() []model.UserData {
	dbSql, _ := data.ConnectDB()
	usersData, _ := data.GetAllUserData(dbSql)
	return usersData
}

func HandlerListPage(w http.ResponseWriter, r *http.Request) {
	usersData := getListUsers()
	tmpl, err := template.ParseFiles("web/templates/list.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		log.Fatalf("Error loading list-template: %v", err)
	}
	if err := tmpl.Execute(w, usersData); err != nil {
		http.Error(w, "Error rendering list of users", http.StatusInternalServerError)
		log.Fatalf("Error rendering list of users: %v", err)
	}
}
