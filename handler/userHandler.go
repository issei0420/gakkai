package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

var userTemps *template.Template

func userParse() error {
	tmps, err := template.ParseFiles("templates/user/index.html", "templates/user/login.html",
		"templates/user/signUp.html", "templates/user/status.html")
	if err != nil {
		return fmt.Errorf("userParse: %v", err)
	}
	userTemps = tmps
	return nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := userParse()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userTemps.ExecuteTemplate(w, "index.html", nil)
}
