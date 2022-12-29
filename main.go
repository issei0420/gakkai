package main

import (
	"fmt"
	"gakkai/handler"
	"net/http"
	"os"
)

func main() {
	fmt.Println("go started")

	dir, _ := os.Getwd()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static/"))))
	http.Handle("/pdf/", http.StripPrefix("/pdf/", http.FileServer(http.Dir(dir+"/pdf/"))))
	// user
	http.HandleFunc("/index", handler.IndexHandler)
	http.HandleFunc("/status", handler.StatusHandler)
	// editor
	http.HandleFunc("/list", handler.ListHandler)
	http.HandleFunc("/register", handler.RegisterHandler)
	// login
	http.HandleFunc("/userLogin", handler.UserLoginHandler)
	http.HandleFunc("/signUp", handler.SignUpHandler)
	http.HandleFunc("/editorLogin", handler.EditorLoginHandler)

	http.ListenAndServe(":8080", nil)
}
