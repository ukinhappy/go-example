package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Open(addr, user, pwd, dbname string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pwd, addr, dbname)+"?charset=utf8mb4&loc=Asia%2FShanghai&parseTime=true")
	return db, err
}
