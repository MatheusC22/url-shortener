package models

type Url struct {
	Url_hash     string
	Url_original string
	Created_at   string
	Expires_at   string
	User_id      string
}

type UrlCreateRequest struct {
	Url_original string `json:"url_original" binding:"required"`
	User_id      string `json:"user_id" binding:"required"`
}
type UrlUpdateRequest struct {
	Url_original string `json:"url_original" binding:"required"`
}
type UrlResponse struct {
	Url_hash     string `json:"url_hash"`
	Url_original string `json:"url_original"`
}
