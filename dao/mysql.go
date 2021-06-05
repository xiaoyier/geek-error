package dao

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {

	var err error
	dsn := "root:root@tcp(localhost)/user?charset=utf8&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil || db == nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Second * 60)
	db.SetMaxIdleConns(30)
	db.SetMaxOpenConns(30)
}

func GetDB() *sql.DB {
	return db
}
