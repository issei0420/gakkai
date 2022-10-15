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

	http.HandleFunc("/index", handler.IndexHandler)
	http.ListenAndServe(":8080", nil)
}
