package main

import (
	"fmt"
	"log"
	"upper.io/db.v3/mysql"
)

type ukinhappytest struct {
	ID     int
	Filed1 int `db:"filed1"`
	Filed2 int `db:"filed2"`
}

func main() {
	mDsn, err := mysql.ParseURL("root:MYSQL123456@tcp(127.0.0.1:3306)/ukinhappy_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db, err := mysql.Open(mDsn)
	if err != nil {
		log.Fatal(err)
	}
	db.SetLogging(true)
	db.SetMaxOpenConns(100)
	var row ukinhappytest
	if err := db.SelectFrom("ukinhappy_test").Where("filed1=?", 2).One(&row); err != nil {
		log.Fatal(err)
	}
	fmt.Println(row)
}
