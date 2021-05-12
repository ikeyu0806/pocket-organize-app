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

type RequestTokenResponse struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

func get_request_token(w http.ResponseWriter) {
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

	body, err := ioutil.ReadAll(res.Body)

	var token_response RequestTokenResponse

	if err := json.Unmarshal(body, &token_response); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		fmt.Println("http request error")
		log.Fatal(err)
	} else {
		fmt.Println("http request success")
		fmt.Println(token_response.Code)
	}
	w.Write(body)
}

func main() {
	fs := http.FileServer(http.Dir("/workspace/pocket_app"))

	http.HandleFunc("/auth_pocket", func(w http.ResponseWriter, r *http.Request) {
		get_request_token(w)
	})

	http.Handle("/", fs)
	http.ListenAndServe(":80", nil)
}
