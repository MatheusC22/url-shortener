package services

import (
	"fmt"
	"goAPI/db"
	"goAPI/models"
)

type userService struct {
}

func NewUserService() *userService {
	return &userService{}
}

func (u *userService) Insert(user models.UserRequest) (user_id string, err error) {
	conn, err := db.OppenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	_, err = conn.Exec(fmt.Sprintf("INSERT INTO  users (username,user_email,user_password) VALUES ('%s','%s','%s')", user.Username, user.User_email, user.User_password))
	if err != nil {
		return
	}
	err = conn.QueryRow(`SELECT @last_uuid as user_id`).Scan(&user_id)

	return
}

func (u *userService) Update(user_id string, user models.UserRequest) (int64, error) {
	conn, err := db.OppenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(fmt.Sprintf("UPDATE users SET username = %s,user_email = %s, user_password = %s WHERE user_id = %s", user.Username, user.User_email, user.User_password, user_id))

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (u *userService) GetById(user_id string) (user models.User, err error) {
	conn, err := db.OppenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(fmt.Sprintf("SELECT user_id,user_email,username,user_password FROM users WHERE user_id = '%s'", user_id))

	err = row.Scan(&user.User_id, &user.User_email, &user.Username, &user.User_password)
	return
}

func (u *userService) GetByEmail(user_email string) (user models.User, err error) {
	conn, err := db.OppenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(fmt.Sprintf("SELECT user_id,user_email,username,user_password FROM users WHERE user_email = '%s'", user_email))

	err = row.Scan(&user.User_id, &user.User_email, &user.Username, &user.User_password)
	return
}

func (u userService) GetAll() (users []models.UserResponse, err error) {
	conn, err := db.OppenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT user_email,username,user_id FROM users`)

	if err != nil {
		return
	}
	for rows.Next() {
		var user models.UserResponse

		err = rows.Scan(&user.User_email, &user.Username, &user.User_id)
		if err != nil {
			continue
		}

		users = append(users, user)
	}

	return users, err
}

func (u *userService) Delete(user_id string) (int64, error) {
	conn, err := db.OppenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(fmt.Sprintf("DELETE FROM users WHERE user_id = '%s'", user_id))

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (u *userService) UserExists(user_email string) (bool, error) {
	conn, err := db.OppenConnection()
	var count int
	if err != nil {
		return false, err
	}
	defer conn.Close()

	err = conn.QueryRow(fmt.Sprintf("SELECT COUNT(user_id) FROM users WHERE user_email = '%s'", user_email)).Scan(&count)

	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, err
	}

	return true, err
}
