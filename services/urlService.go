package services

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"goAPI/db"
	"goAPI/models"
)

type urlService struct {
}

func NewUrlService() *urlService {
	return &urlService{}
}

func (u *urlService) Insert(url models.UrlDTO) (url_hash string, err error) {
	conn, err := db.OppenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	bytes := make([]byte, 3)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	url_hash = hex.EncodeToString(bytes)

	_, err = conn.Exec(fmt.Sprintf("INSERT INTO urls (url_hash,url_original,user_id) VALUES ('%s','%s','%s')", url_hash, url.Url_original, url.User_id))
	if err != nil {
		return
	}
	return
}

func (u *urlService) Get(url_hash string) (url models.UrlDTO, err error) {
	conn, err := db.OppenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(fmt.Sprintf("SELECT url_hash,url_original,useer_id FROM urls WHERE url_hash = %s", url_hash))

	err = row.Scan(&url.Url_hash, &url.Url_original, &url.User_id)

	return
}
func (u *urlService) GetAll() (urls []models.UrlDTO, err error) {
	conn, err := db.OppenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM urls`)

	if err != nil {
		return
	}
	for rows.Next() {
		var url models.UrlDTO

		err = rows.Scan(&url.Url_hash, &url.Url_original, &url.User_id)
		if err != nil {
			continue
		}

		urls = append(urls, url)
	}

	return
}

func (u *urlService) Delete(url_hash string) (int64, error) {
	conn, err := db.OppenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(fmt.Sprintf("DELETE FROM urls WHERE url_hash = %s", url_hash))

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (u *urlService) Update(url_hash string, url models.UrlDTO) (int64, error) {
	conn, err := db.OppenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(fmt.Sprintf("UPDATE urls SET url_original = %s WHERE url_hash = %s", url.Url_original, url_hash))

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
