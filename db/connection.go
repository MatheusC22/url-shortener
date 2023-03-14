package db

import (
	"database/sql"
	"fmt"
	"goAPI/configs"

	_ "github.com/go-sql-driver/mysql"
)

func OppenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)

	conn, err := sql.Open("mysql", sc)

	if err != nil {
		panic(err)
	}
	err = conn.Ping()

	return conn, err
}

func Con() string {
	conf := configs.GetDB()

	sc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
	return sc
}
