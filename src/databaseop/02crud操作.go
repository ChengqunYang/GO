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

// 查询单个记录
func queryOne(id int) {
	var u1 user
	// 写查询单条记录的sql语句
	sqlStr := `select id,name,age from user where id = ?;`
	// 执行
	rowObj := db.QueryRow(sqlStr, 2) // 从连接池里拿出一个连接出来去数据库查询单条记录
	// 拿到结果
	rowObj.Scan(&u1.id, &u1.name, &u1.age) // scan方法中会释放连接，归还到连接池里面
	// 打印结果
	fmt.Printf("u1:%#v\n", u1)
}

// 查询多个记录
func queryMore(n int) {
	//1. 写查询多条的sql语句
	sqlStr := `select id, name, age from user where id > ?;`
	//2. 执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Print("exec %c query failed, err:%v\n", sqlStr, err)
		return
	}
	//3. 一定要关闭rows
	defer rows.Close()
	//4. 循环取值
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
		}
		fmt.Printf("u1:%#v\n", u1)
	}
}
func insert() {
	//1. 写sql语句
	sqlStr := `insert into user(name,age)values("田七",26);`
	//2. 执行
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	// 如果是插入数据的操作，可以拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Println("id:", id)

}
func updateRow(newAge int, id int) {
	sqlStr := `update user set age= ? where id = ?;`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Print("update failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Printf("更新了%d行数据", n)
}
func deleteRow(id int) {
	sqlStr := `delete from user where id= ?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Printf("删除了%d行数据\n", n)
}

/*
	预处理的方式插入多条数据
*/
func prepareInsert() {
	sqlStr := `insert into user(name,age)values(?,?)`
	stmt, err := db.Prepare(sqlStr) // 把sql语句现发给mysql预处理一下
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
		return
	}
	defer stmt.Close()
	// 后续只要拿到stmt去执行一些操作
	//stmt.Exec("王麻子",80)
	//stmt.Exec("小爱同学",5)
	var m = map[string]int{
		"光头强": 30,
		"熊大":  10,
		"熊二":  9,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err %v\n", err)
		return
	}
	fmt.Println("连接数据库成功！")

	//queryOne(2)
	//queryMore(2)
	//insert()
	//updateRow(9000,5)
	//deleteRow(7)
	prepareInsert()

}
