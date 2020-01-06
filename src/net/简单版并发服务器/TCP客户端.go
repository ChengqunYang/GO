package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("net Dial err = ", err)
		return
	}

	//main调用完毕,关闭连接
	defer conn.Close()

	go func() {
		//从键盘输入内容,发送给服务器
		str := make([]byte, 1024)
		for true {
			n, err1 := os.Stdin.Read(str) //从键盘读取内容,存入str
			if err != nil {
				fmt.Println("os.Stdin.Read err = ", err1)
				return
			}
			//把输入的内容给服务器发送
			conn.Write(str[:n])
		}
	}()

	//接收服务器回复的数据,新任务
	//定义一个切片缓存
	buf := make([]byte, 1024)
	for true {
		n, err := conn.Read(buf) //接收服务器的请求
		if err != nil {
			fmt.Println("conn.Read err = ", err)
			return
		}
		fmt.Println(string(buf[:n])) //打印接收到的内容
	}
}
