package main

import (
	"net/http"
	"text/template"
)

type Page struct {
	Title   string
	Message string
}

func handler(w http.ResponseWriter, r *http.Request) {
	page := Page{}
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
