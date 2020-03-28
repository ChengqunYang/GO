package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //没有使用，那么调用init()方法
)

/*Go语言中的database/sql包提供了保证SQL或类SQL数据库的泛用接口，
并不提供具体的数据库驱动。使用database/sql包时必须注入（至少）一个数据库驱动。
*/
func main() {
	//数据库信息
	dsn := "root:975864@tcp(127.0.0.1:3306)/gotest"
	//连接数据库
	db, err := sql.Open("mysql", dsn) //不会校验用户名和密码是否正确
	if err != nil {                   // dsn的格式不正确时会报错
		fmt.Printf("dsn %s invalid,err:%v\n", dsn, err)
		return
	}
	err = db.Ping() // 尝试连接数据库
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功!")
}
