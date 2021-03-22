package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)
func main() {
	query()
	return
	fmt.Println("获取所有值")
	allresult()
	fmt.Println("获取服务器版本号")
	version()
	fmt.Println("获取表名字")
	tablename()
	fmt.Println("获取数据库名字")
	dbname()
}



func Open(addr, user, pwd, dbname string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pwd, addr, dbname)+"?charset=utf8mb4&loc=Asia%2FShanghai&parseTime=true")
	return db, err
}

func query() {
	db, err := Open("127.0.0.1", "root", "MYSQL123456", "ukinhappy_test")
	if err != nil {
		log.Fatal(err)

	}
	// 查询所有的变量
	sql := "select id from `test` where value = ?"
	r := db.QueryRow(sql, "8")
	var value int
	for {
		if err:=r.Scan(&value); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(value)
	}

}
func allresult() {
	db, err := Open("127.0.0.1", "root", "MYSQL123456", "ukinhappy_test")
	if err != nil {
		log.Fatal(err)

	}
	// 查询所有的变量
	sql := "select value from `test` where id = " + "1 or 1 = 1 "
	r, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	var value int
	for r.Next() {
		r.Scan(&value)
		fmt.Println(value)
	}
	defer r.Close()

	// 防止sql注入预处理
	sql = "select value from `test` where id = ?"
	stm, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stm.Query("1 or 1 = 1 ")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&value)
		fmt.Println(value)
	}
	defer rows.Close()

}

func version() {
	db, err := Open("127.0.0.1", "root", "MYSQL123456", "ukinhappy_test")
	if err != nil {
		log.Fatal(err)

	}
	// 查询所有的变量
	sql := "select value from `test` where id = " + "1 union select version() as value  "
	r, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	var value string
	for r.Next() {
		r.Scan(&value)
		fmt.Println(value)
	}
	defer r.Close()

	// 防止sql注入预处理
	sql = "select value from `test` where id = ?"
	stm, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stm.Query(" 1 union select version() as value ")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&value)
		fmt.Println(value)
	}
	defer rows.Close()

}

func tablename() {
	db, err := Open("127.0.0.1", "root", "MYSQL123456", "ukinhappy_test")
	if err != nil {
		log.Fatal(err)

	}
	// 查询所有的变量
	sql := "select value from `test` where id = " + "1 union select database() as value  "
	r, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	var value string
	for r.Next() {
		r.Scan(&value)
		fmt.Println(value)
	}
	defer r.Close()

	// 防止sql注入预处理
	sql = "select value from `test` where id = ?"
	stm, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stm.Query(" 1 union select database() as value ")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&value)
		fmt.Println(value)
	}
	defer rows.Close()

}

func dbname() {
	db, err := Open("127.0.0.1", "root", "MYSQL123456", "ukinhappy_test")
	if err != nil {
		log.Fatal(err)

	}
	// 查询所有的变量
	sql := "select value from `test` where id = " + "1 union SELECT GROUP_CONCAT(schema_name) as value FROM information_schema.schemata   "
	r, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	var value string
	for r.Next() {
		r.Scan(&value)
		fmt.Println(value)
	}
	defer r.Close()

	// 防止sql注入预处理
	sql = "select value from `test` where id = ?"
	stm, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stm.Query(" 1 union SELECT GROUP_CONCAT(schema_name) as value FROM information_schema.schemata  ")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&value)
		fmt.Println(value)
	}
	defer rows.Close()

}
