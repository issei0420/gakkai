package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var temp = template.Must(template.ParseFiles("templates/user/index.html"))

func hello(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", nil)
}

func main() {
	fmt.Println("go started")

	dir, _ := os.Getwd()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static/"))))
	http.Handle("/pdf/", http.StripPrefix("/pdf/", http.FileServer(http.Dir(dir+"/pdf/"))))

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
