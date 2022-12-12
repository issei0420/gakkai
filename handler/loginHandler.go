package handler

import (
	"html/template"
	"net/http"
)

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	tf, err := template.ParseFiles("templates/user/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tf.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	tf, err := template.ParseFiles("templates/user/signUp.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tf.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
