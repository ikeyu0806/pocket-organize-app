package entity

type ArticleRequestParam struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
	Count       string `json:"count"`
	Tag         string `json:"tag"`
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
