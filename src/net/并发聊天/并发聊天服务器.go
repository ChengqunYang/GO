package main

import (
	"fmt"
	"net"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	//获取客户端网络地址
	cliAddr := conn.RemoteAddr().String()
	//创建一个结构体,用户名和地址默认一样
	cli := Client{make(chan string), cliAddr, cliAddr}
	//把结构体添加到map
	onlineMap[cliAddr] = cli

	//新开一个协程,专门给当前客户端发送信息
	go WriteMsgToClint(cli, conn)

	//广播某个人在线
	message <- "{" + cli.Addr + "}" + cli.Name + ":  longin"
	for true {

	}
}

func WriteMsgToClint(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func Manager() {
	//给map分配空间
	onlineMap = make(map[string]Client)
	for true {
		msg := <-message //没有消息前,阻塞
		//遍历map,给map中的每个成员都发送此消息
		for _, client := range onlineMap {
			client.C <- msg
		}
	}
}

type Client struct {
	C    chan string //用户发送数据的管道
	Name string      //用户名
	Addr string      //网络地址
}

//保存在线用户
var onlineMap map[string]Client

//创建一个通信的管道
var message = make(chan string)

func main() {
	//监听
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	defer listener.Close()

	//新开一个协程,用来转发消息,只要有消息来,就遍历map,给每个成员都发送这个消息
	go Manager()

	//主协程,循环阻塞,等待用户连接
	for true {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err = ", err)
			continue
		}
		go HandleConn(conn) //处理用户连接
	}
}
