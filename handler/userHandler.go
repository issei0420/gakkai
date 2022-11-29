package handler

import (
	"fmt"
	"html/template"
	"log"
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
	if err = userTemps.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("############")
	if err := userTemps.ExecuteTemplate(w, "status.html", nil); err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
