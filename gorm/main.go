package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:MYSQL123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	value := new(struct {
		Book  string
		Price int
	})
	if err := db.Table("book").Where("book =?", "go-zero").First(&value).Error; err != nil {

		log.Fatal(err)
	}
	fmt.Println(value)

	value = new(struct {
		Book  string
		Price int
	})

	if err := db.Table("book").Where("book= 1 or 1 = 1 ").First(&value).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
}
