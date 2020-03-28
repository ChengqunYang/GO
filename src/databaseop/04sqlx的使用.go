package main

import (
	//"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB // 连接池

func initDB() (err error) {
	dsn := "root:975864@tcp(127.0.0.1:3306)/gotest"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

type user struct {
	Id   int
	Name string
	Age  int
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
		return
	}
	sqlStr := `select id,name,age from user where id = 1;`
	var u user
	db.Get(&u, sqlStr)
	fmt.Printf("u:%#v\n", u)

	var userList []user
	sqlStr2 := `select * from user;`
	db.Select(&userList, sqlStr2)
	fmt.Println(userList)
}
