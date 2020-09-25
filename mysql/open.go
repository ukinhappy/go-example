package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1)/ukinhappy_test?charset=utf8mb4&loc=Asia%2FShanghai&parseTime=true")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.DB().Ping()
}
