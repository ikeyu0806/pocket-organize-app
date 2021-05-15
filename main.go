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

type AccessTokenParam struct {
	ConsumerKey string `json:"consumer_key"`
	Code        string `json:"code"`
}

type PocketRequestParam struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
}

type RequestTokenResponse struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	UserName    string `json:"username"`
}

type ArticleResponse struct {
	Status       int `json:"status"`
	Complete     int `json:"complete"`
	ArticleDatas []interface{}
}

type ArticleData struct {
	ItemID        string `json:"item_id"`
	ResolvedID    string `json:"resolved_id"`
	GivenURL      string `json:"given_url"`
	GivenTitle    string `json:"given_title"`
	Favorite      string `json:"favorite"`
	Status        string `json:"status"`
	TimeAdded     string `json:"time_added"`
	TimeUpdated   string `json:"time_updated"`
	TimeRead      string `json:"time_read"`
	TimeFavorited string `json:"time_favorited"`
	SortID        int    `json:"sort_id"`
	ResolvedTitle string `json:"resolved_title"`
	ResolvedURL   string `json:"resolved_url"`
	Excerpt       string `json:"excerpt"`
	IsArticle     string `json:"is_article"`
	IsIndex       string `json:"is_index"`
	HasVideo      string `json:"has_video"`
	HasImage      string `json:"has_image"`
	WordCount     string `json:"word_count"`
	Lang          string `json:"lang"`
	TopImageURL   string `json:"top_image_url"`
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

	if err != nil {
		fmt.Println("http response error")
		log.Fatal(err)
	}

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

func get_access_token(w http.ResponseWriter, r *http.Request) {
	access_token_param := new(AccessTokenParam)
	access_token_param.ConsumerKey = os.Getenv(("POCKET_COSUMER_KEY"))
	access_token_param.Code = r.FormValue("code")

	access_token_json, _ := json.Marshal(access_token_param)
	fmt.Printf("[access_token] %s\n", string(access_token_json))

	req, err := http.NewRequest("POST", "https://getpocket.com/v3/oauth/authorize", bytes.NewBuffer(access_token_json))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Accept", "application/json")

	client := &http.Client{Timeout: time.Duration(180) * time.Second}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("http response error")
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	var token_response AccessTokenResponse

	if err := json.Unmarshal(body, &token_response); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		fmt.Println("http request error")
		log.Fatal(err)
	} else {
		fmt.Println("http request success")
		fmt.Println(token_response.AccessToken)
	}
	w.Write(body)
}

func get_articles(w http.ResponseWriter, r *http.Request) {
	pocket_request_param := new(PocketRequestParam)
	pocket_request_param.ConsumerKey = os.Getenv(("POCKET_COSUMER_KEY"))
	pocket_request_param.AccessToken = r.FormValue("access_token")

	pocket_request_json, _ := json.Marshal(pocket_request_param)
	fmt.Printf("[pocket_request_json] %s\n", string(pocket_request_json))

	req, err := http.NewRequest("POST", "https://getpocket.com/v3/get", bytes.NewBuffer(pocket_request_json))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Accept", "application/json")

	client := &http.Client{Timeout: time.Duration(180) * time.Second}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("http response error")
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	var article_response ArticleResponse

	if err := json.Unmarshal(body, &article_response); err != nil {
		fmt.Println("Unmarshal err")
		log.Fatal(err)
	}
	fmt.Println("article_response")
	fmt.Println(article_response)

	for _, data := range article_response.ArticleDatas {
		fmt.Println("data")
		fmt.Println(data)
	}

	if err != nil {
		fmt.Println("http request error")
		log.Fatal(err)
	} else {
		fmt.Println("http request success")
		fmt.Println(article_response)
	}
	w.Write(body)
}

func main() {
	fmt.Println("start server")
	fs := http.FileServer(http.Dir("/workspace/pocket_app"))

	http.HandleFunc("/get_request_token", func(w http.ResponseWriter, r *http.Request) {
		get_request_token(w)
	})

	http.HandleFunc("/get_access_token", func(w http.ResponseWriter, r *http.Request) {
		get_access_token(w, r)
	})

	http.HandleFunc("/get_articles", func(w http.ResponseWriter, r *http.Request) {
		get_articles(w, r)
	})

	http.Handle("/", fs)
	http.ListenAndServe(":80", nil)
}
