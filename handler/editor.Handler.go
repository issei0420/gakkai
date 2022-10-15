package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

var editorTemps *template.Template

func editorParse() error {
	tmps, err := template.ParseFiles("templates/editor/list.html", "templates/editor/register.html")
	if err != nil {
		return fmt.Errorf("editorParse: %v", err)
	}
	editorTemps = tmps
	return nil
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	err := editorTemps.ExecuteTemplate(w, "list.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	err := editorParse()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = editorTemps.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
