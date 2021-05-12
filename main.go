package main

import (
	"net/http"
)

type Page struct {
	Title   string
	Message string
}

func main() {
	fs := http.FileServer(http.Dir("/workspace/pocket_app"))
	http.Handle("/", fs)
	http.ListenAndServe(":80", nil)
}
