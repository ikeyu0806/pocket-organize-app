package main

import (
	"fmt"
	"net/http"
	"pocket-organize-app/interfaces"
)

func main() {
	fmt.Println("start server")
	fs := http.FileServer(http.Dir("/usr/local/go/src/pocket-organize-app"))

	http.HandleFunc("/get_request_token", func(w http.ResponseWriter, r *http.Request) {
		interfaces.Get_request_token(w)
	})

	http.HandleFunc("/get_access_token", func(w http.ResponseWriter, r *http.Request) {
		interfaces.Get_access_token(w, r)
	})

	http.HandleFunc("/get_articles", func(w http.ResponseWriter, r *http.Request) {
		interfaces.Get_articles(w, r)
	})

	http.Handle("/", fs)
	http.ListenAndServe(":80", nil)
}
