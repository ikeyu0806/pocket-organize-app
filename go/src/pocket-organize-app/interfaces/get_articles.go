package interfaces

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"pocket-organize-app/entity"
)

func Get_articles(w http.ResponseWriter, r *http.Request) {
	pocket_request_param := new(entity.ArticleRequestParam)
	pocket_request_param.ConsumerKey = os.Getenv(("POCKET_COSUMER_KEY"))
	pocket_request_param.AccessToken = r.FormValue("access_token")
	pocket_request_param.Count = 5

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

	var article_response entity.ArticleResponse

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
