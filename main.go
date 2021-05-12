package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type RequestTokenParam struct {
	ConsumerKey string `json:"consumer_key"`
	RedirectURL string `json:"redirect_uri"`
}

func get_request_token() {
	request_token_param := new(RequestTokenParam)
	request_token_param.ConsumerKey = os.Getenv(("POCKET_COSUMER_KEY"))
	request_token_param.RedirectURL = os.Getenv(("REDIRECT_URI"))

	request_token_json, _ := json.Marshal(request_token_param)
	fmt.Printf("[+] %s\n", string(request_token_json))

	req, err := http.NewRequest("POST", "https://getpocket.com/v3/oauth/request", bytes.NewBuffer(request_token_json))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Accept", "application/json")

	client := &http.Client{Timeout: time.Duration(180) * time.Second}
	res, err := client.Do(req)

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("http request error")
		log.Fatal(err)
	} else {
		fmt.Println("http request success")
		fmt.Println(string(b))
	}
}

func main() {
	fs := http.FileServer(http.Dir("/workspace/pocket_app"))

	http.HandleFunc("/auth_pocket", func(w http.ResponseWriter, r *http.Request) {
		get_request_token()
	})

	http.Handle("/", fs)
	http.ListenAndServe(":80", nil)
}
