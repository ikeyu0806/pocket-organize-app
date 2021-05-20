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

func AddTags(w http.ResponseWriter, r *http.Request) {
	request_body, err := ioutil.ReadAll(r.Body)
	fmt.Println("Body")
	fmt.Println(string(request_body))

	var request_param entity.ClientModifyRequestParam

	if err := json.Unmarshal(request_body, &request_param); err != nil {
		fmt.Println("response Unmarshal err")
		log.Fatal(err)
	}

	modify_action := new(entity.ModifyAction)
	modify_action.Action = "tags_add"
	modify_action.ItemID = request_param.ItemIDs
	modify_action.Tags = request_param.Tags
	fmt.Println("request_param.Tags")
	fmt.Println(request_param.Tags)

	modify_request_param := new(entity.ModifyRequestParam)
	modify_request_param.ConsumerKey = os.Getenv(("POCKET_COSUMER_KEY"))
	modify_request_param.AccessToken = request_param.AccessToken
	modify_request_param.Actions = append(modify_request_param.Actions, modify_action)
	modify_request_json, _ := json.Marshal(modify_request_param)

	fmt.Printf("[tags add request_json] %s\n", string(modify_request_json))

	req, err := http.NewRequest("POST", "https://getpocket.com/v3/send", bytes.NewBuffer(modify_request_json))

	if err != nil {
		fmt.Println("pocket modify request error")
	}

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

	var response entity.ActionResult

	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("response Unmarshal err")
		log.Fatal(err)
	}
	fmt.Println("response")
	fmt.Println(response)

	if err != nil {
		fmt.Println("http request error")
		log.Fatal(err)
	} else {
		fmt.Println("http request success")
		fmt.Println(response)
	}
	w.Write(body)
}
