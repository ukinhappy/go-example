package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
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

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:MYSQL123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{})

	u1:=&User{Name: "n1",Email: "e1"}
	result := db.Create(u1)
	fmt.Println(u1)
	fmt.Println(result.Error, result.RowsAffected)


	//批量插入
	var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	db.Create(&users)
	for _, user := range users {
		fmt.Println(user.ID)
	}

	//
	db.Model(u1).Update("name","N1")
	db.Model(&User{}).Update("name","N1")

	//查询
	result =db.Last(u1)
	fmt.Println(u1)
}
