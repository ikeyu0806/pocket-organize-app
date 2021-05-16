package entity

type RequestTokenParam struct {
	ConsumerKey string `json:"consumer_key"`
	RedirectURL string `json:"redirect_uri"`
}

type RequestTokenResponse struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

type AccessTokenParam struct {
	ConsumerKey string `json:"consumer_key"`
	Code        string `json:"code"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	UserName    string `json:"username"`
}
