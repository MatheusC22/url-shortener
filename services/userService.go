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

func (u *userService) Insert(user models.UserDTO) (user_id string, err error) {
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

func (u *userService) Update(user_id string, user models.UserDTO) (int64, error) {
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

func (u *userService) Get(user_id string) (user models.User, err error) {
	conn, err := db.OppenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(fmt.Sprintf("SELECT * FROM users WHERE user_id = %s", user_id))

	err = row.Scan(&user.User_id, &user.User_email, &user.User_password, &user.Username)
	return
}

func (u userService) GetAll() (users []models.User, err error) {
	conn, err := db.OppenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM users`)

	if err != nil {
		return
	}
	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.User_id, &user.User_email, &user.User_password, &user.Username)
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

	res, err := conn.Exec(fmt.Sprintf("DELETE FROM users WHERE id = %s", user_id))

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
