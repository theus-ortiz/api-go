package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	dsn := "root@tcp(localhost:3306)/user_api?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	if err := db.Ping(); err != nil {
		panic(err.Error())
	}

	return db
}
