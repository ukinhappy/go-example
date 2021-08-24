package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	gorm.Model
	Name         string
	Email        string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
}

var db *gorm.DB

type ukinhappytest struct {
	ID     int
	Filed1 int `gorm:"column:filed1"`
	Filed2 int `gorm:"column:filed2"`
}

func main() {
	var err error
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:MYSQL123456@tcp(127.0.0.1:3306)/ukinhappy_test?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("open", err)
	}

	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)

	for i := 0; i < 4; i++ {
		go func() {
			if err := db.Table("ukinhappy_test").Where("id=?", 1).
				Update("filed1", 1).
				Update("filed2", 1).Error; err != nil {
				log.Fatal("update 1 ", err)
			}
		}()
	}

	for i := 0; i < 3; i++ {
		go func() {
			if err := db.Table("ukinhappy_test").Where("id=?", 1).
				Update("filed1", 2).
				Update("filed2", 2).Error; err != nil {
				log.Fatal("update 2 ", err)
			}
		}()
	}

	for i := 0; i < 3; i++ {
		var value ukinhappytest
		if err := db.Table("ukinhappy_test").Where("id=?", 1).First(&value).Error; err != nil {
			log.Fatal("query", err.Error)
		}
		if value.Filed1 != value.Filed2 {
			fmt.Println(value)
		}


	}

	time.Sleep(time.Minute)
}
