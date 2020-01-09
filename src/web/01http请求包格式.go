package main

import (
	"fmt"
	"net"
)

func main() {
	//监听
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}
	defer listener.Close()
	//阻塞等待用户的连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err", err)
		return
	}
	defer conn.Close()

	//接收客户端的数据
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if n == 0 {
		fmt.Println("read err = ", err)
		return
	}
	fmt.Printf("#%v#", string(buf[:n]))

}
