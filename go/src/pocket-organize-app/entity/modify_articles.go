package entity

type ClientModifyRequestParam struct {
	AccessToken string   `json:"access_token"`
	Tags        []string `json:"tags"`
	// TODO配列にする
	ItemIDs string `json:"item_ids"`
}

type ModifyRequestParam struct {
	ConsumerKey string          `json:"consumer_key"`
	AccessToken string          `json:"access_token"`
	Actions     []*ModifyAction `json:"actions"`
}

type ModifyAction struct {
	Action string   `json:"action"`
	ItemID string   `json:"item_id"`
	Tags   []string `json:"tags"`
}

type ActionResult struct {
	Status int `json:"status"`
}
