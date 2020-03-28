package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //没有使用，那么调用init()方法
)

var db *sql.DB // 连接池

func initDB() (err error) {
	dsn := "root:975864@tcp(127.0.0.1:3306)/gotest"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping() // 尝试连接数据库
	if err != nil {
		return
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10) // 当池子里的连接全部在使用时会阻塞
	db.SetMaxIdleConns(5)  // 设置连接池中最大的闲置连接数
	return
}

type user struct {
	id   int
	name string
	age  int
}

func transactionDemo() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed,err:%v\n", err)
		return
	}
	//执行多个sql操作
	sqlStr := `update user set age = ? where id = ?;`
	stmt, err := tx.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
		return
	}
	_, err = stmt.Exec(50, 8)
	if err != nil {
		//要回滚
		tx.Rollback()
		fmt.Println("执行sqlstr1出错了,要回滚!")
		return
	}
	_, err = stmt.Exec("50")
	if err != nil {
		//要回滚
		tx.Rollback()
		fmt.Println("执行sqlstr2出错了,要回滚!")
		return
	}
	//上面两步都执行成功了
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("提交出错了,要回滚!")
		return
	}

}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err %v\n", err)
	}
	transactionDemo()
}
