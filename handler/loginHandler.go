package handler

import (
	"encoding/json"
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

type Data struct {
	Mail     string
	Password string
	UserType string
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var d Data
		json.NewDecoder(r.Body).Decode(&d)

		type response struct {
			Invalid int `json:"invalid"`
			// ResMap  map[string]int `json:"resmap"`
		}

		res := response{Invalid: 0}
		b, err := json.MarshalIndent(res, "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(b)

	} else {
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
}

func EditorLoginHandler(w http.ResponseWriter, r *http.Request) {
	tf, err := template.ParseFiles("templates/editor/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tf.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
