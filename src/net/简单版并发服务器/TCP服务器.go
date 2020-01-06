package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	//函数调用完毕,自动关闭conn
	defer conn.Close()
	addr := conn.RemoteAddr()
	fmt.Println(addr, "connect sucess")

	buf := make([]byte, 2048)
	for true {
		//读取用户数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Printf("[%s]:%s\n", addr, string(buf[:n-1]))
		if "exit" == string(buf[:n-1]) {
			fmt.Println(addr, "退出了")
			return
		}
		//把数据转化为大写,并发送给用户
		conn.Write([]byte(strings.ToUpper(string(buf[:n-1]))))
	}

}

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listener.Close()

	//接收多个用户
	for true {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err = ", err)
			return
		}

		//处理用户请求,每来一个用户新建一个协程
		go HandleConn(conn)
	}
}
