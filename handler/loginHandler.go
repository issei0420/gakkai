package handler

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"gakkai/db"
	"gakkai/lib"
	"html/template"
	"log"
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
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var d Data
		json.NewDecoder(r.Body).Decode(&d)

		isUnique, err := db.IsUniqueParticipant(d.Mail)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		type response struct {
			Valid int `json:"valid"`
		}
		var res response

		if isUnique {
			pHash := lib.MakeHash(d.Password)
			err = db.InsertParticipant(d.Mail, pHash)
			if err != nil {
				fmt.Println(pHash)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Fatal(err)
				return
			}
			res = response{Valid: 1}
		} else {
			res = response{Valid: 0}
		}

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
