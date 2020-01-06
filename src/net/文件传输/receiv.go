package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func receivFile(fileName string, conn net.Conn) {
	//新建文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Creat err = ", err)
		return
	}
	//接收内容,接收多少写多少
	buf := make([]byte, 4*1024)
	for true {
		n, err := conn.Read(buf) //接收对放发送的内容
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err = ", err)
			}
			return
		}
		if n == 0 {
			fmt.Println("文件接收完毕")
			return
		}
		f.Write(buf[:n]) //将读取到的内容写入文件

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

	//阻塞等待用户连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err = ", err)
		return
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}
	//把读取到的文件路径保存
	fileName := string(buf[:n])

	//给对方回复ok
	conn.Write([]byte("ok"))

	//接受文件内容
	receivFile(fileName, conn)
}
