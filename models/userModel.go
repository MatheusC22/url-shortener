package models

type User struct {
	User_id       string `json:"user_id"`
	Username      string `json:"username"`
	User_email    string `json:"user_email"`
	User_password string `json:"user_password"`
}
