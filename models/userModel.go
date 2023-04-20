package models

type User struct {
	User_id       string
	Username      string
	User_email    string
	User_password string
}

type UserRequest struct {
	Username      string `json:"username" binding:"required"`
	User_email    string `json:"user_email" binding:"required"`
	User_password string `json:"user_password" binding:"required"`
}

type UserResponse struct {
	Username   string `json:"username"`
	User_email string `json:"user_email"`
	User_id    string `json:"user_id"`
}

type UserJWTPayload struct {
	User_id string `json:"user_id"`
}

type UserLoginRequest struct {
	User_email    string `json:"user_email" binding:"required"`
	User_password string `json:"user_password" binding:"required"`
}
