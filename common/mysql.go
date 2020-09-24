package common

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlConn() (db *sql.DB, err error) {
	msn := "root:joeytest123...@tcp(localhost:3306)/imooc?charset=utf8"
	db, err = sql.Open("mysql", msn)
	if err != nil {
		panic(err)
	}
	return db, err
}
