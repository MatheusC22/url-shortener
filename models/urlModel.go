package models

type Url struct {
	Url_hash     string
	Url_original string
	Created_at   string
	Expires_at   string
	User_id      string
}

type UrlDTO struct {
	Url_hash     string `json:"url_hash"`
	Url_original string `json:"url_original"`
	User_id      string `json:"user_id"`
}
