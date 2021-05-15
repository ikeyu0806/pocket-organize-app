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

func Get_access_token(w http.ResponseWriter, r *http.Request) {
	access_token_param := new(entity.AccessTokenParam)
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

	var token_response entity.AccessTokenResponse

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
